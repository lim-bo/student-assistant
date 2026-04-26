package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/student-assistant/internal/model"
	eventservice "github.com/student-assistant/internal/service/event_service"
	shedule_service "github.com/student-assistant/internal/service/schedule_service"
	"golang.org/x/sync/errgroup"
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

type PlanResponse struct {
	Timeline []*model.TimelineItem `json:"timeline"`
}

func (h *Handler) Plan(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse(time.DateOnly, r.URL.Query().Get("date"))
	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
		slog.Error("Provided invalid date query param")
		return
	}

	group := r.URL.Query().Get("group")
	if group == "" {
		http.Error(w, "Group param required", http.StatusBadRequest)
		slog.Error("Group param not provided")
		return
	}

	eg := errgroup.Group{}

	lessons := make([]model.ScheduleLesson, 0, 4)
	eg.Go(func() error {
		lessons, err = h.etuClient.GetSchedule(group, date)
		if err != nil {
			slog.Error("Error getting lessons", slog.String("error", err.Error()))
			return err
		}
		return nil
	})

	events := make([]*model.Event, 0, 5)
	eg.Go(func() error {
		for _, svc := range h.eventServices {
			eventsChunk, err := svc.GetEvents(eventservice.Options{Date: date, Loc: eventservice.Spb})
			if err != nil {
				slog.Error("Error getting events", slog.String("error", err.Error()))
				return err
			}
			events = append(events, eventsChunk...)
		}
		return nil
	})

	err = eg.Wait()
	if err != nil {
		slog.Error("Error getting data", slog.String("error", err.Error()))
		http.Error(w, "Getting schedule error", http.StatusBadRequest)
		return
	}

	timeline := make([]*model.TimelineItem, 0, 4)
	j := 0
	for i := 0; i < len(lessons); i++ {
		timeline = append(timeline, lessons[i].ToTimelineItem())
		if i == len(lessons)-1 {
			for j < len(events) {
				timeline = append(timeline, events[j].ToTimelineItem())
				j++
			}
		} else {
			for j < len(events) && events[j].StartTime.Before(lessons[i].StartTime) {
				timeline = append(timeline, events[j].ToTimelineItem())
				j++
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PlanResponse{Timeline: timeline})
	slog.Info("Plan provided")
}
