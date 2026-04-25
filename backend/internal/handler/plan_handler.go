package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/student-assistant/internal/model"
	eventservice "github.com/student-assistant/internal/service/event_service"
	shedule_service "github.com/student-assistant/internal/service/schedule_service"
)

type Handler struct {
	etuClient     *shedule_service.EtuClient
	eventServices []eventservice.EventService
}

func NewHandler(etuClient *shedule_service.EtuClient, eventServices []eventservice.EventService) *Handler {
	return &Handler{
		etuClient:     etuClient,
		eventServices: eventServices,
	}
}

type PlanRequest struct {
	Group      string     `json:"group"`
	Faculty    string     `json:"faculty"`
	Date       string     `json:"date"`
	Interests  []int      `json:"interests"`
	HomeCoords [2]float64 `json:"home_coords"`
}

type TimelineItem struct {
	TimeStart string      `json:"time_start"`
	TimeEnd   string      `json:"time_end"`
	Type      string      `json:"type"`
	Title     string      `json:"title"`
	Location  *string     `json:"location,omitempty"`
	Coords    *[2]float64 `json:"coords,omitempty"`
}

type PlanResponse struct {
	Timeline []TimelineItem `json:"timeline"`
}

func (h *Handler) HandlerPlan(w http.ResponseWriter, r *http.Request) {
	var req PlanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
		return
	}

	timeline := make([]TimelineItem, 0)

	lessons, err := h.etuClient.GetSchedule(req.Group, date)
	if err != nil {
		http.Error(w, "schedule error", http.StatusInternalServerError)
		return
	}
	for _, l := range lessons {
		timeline = append(timeline, TimelineItem{
			TimeStart: l.StartTime.Format("15:04"),
			TimeEnd:   l.EndTime.Format("15:04"),
			Type:      l.Type,
			Title:     l.Title,
			Location:  l.Location,
			Coords:    coordsFromLesson(l),
		})
	}

	for _, svc := range h.eventServices {
		events, err := svc.GetEvents(eventservice.Options{Date: date})
		if err != nil {
			continue
		}
		for _, e := range events {
			timeline = append(timeline, TimelineItem{
				TimeStart: e.StartTime,
				TimeEnd:   e.EndTime,
				Type:      "event",
				Title:     e.Title,
				Location:  &e.Location,
				Coords:    &e.Coords,
			})
		}
	}

	// Сортируем по времени начала
	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].TimeStart < timeline[j].TimeStart
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PlanResponse{Timeline: timeline})
}

func coordsFromLesson(l model.ScheduleLesson) *[2]float64 {
	if l.LocationLat == nil || l.LocationLon == nil {
		return nil
	}
	coords := [2]float64{*l.LocationLat, *l.LocationLon}
	return &coords
}
