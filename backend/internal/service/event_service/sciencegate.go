package eventservice

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bytedance/sonic"
	"github.com/student-assistant/internal/config"
	"github.com/student-assistant/internal/model"
	geocoding "github.com/student-assistant/pkg/GeoCoding"
)

const (
	eventsEndpoint = "/event"
)

type ScienceGateService struct {
	geocoder *geocoding.GeoCoder
}

func NewScienceGate(geocoder *geocoding.GeoCoder) *ScienceGateService {
	return &ScienceGateService{
		geocoder: geocoder,
	}
}

func (s *ScienceGateService) GetEvents(opts Options) ([]*model.Event, error) {
	url := config.ScienceGateAPIBaseUrl + eventsEndpoint
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error getting science gate events: %w", err)
	}

	q := req.URL.Query()
	q.Add("periodsAfter", opts.Date.Add(-time.Hour*24).Format(time.DateOnly))
	q.Add("periodsBefore", opts.Date.Add(time.Hour*24).Format(time.DateOnly))
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error getting science gate events: %w", err)
	}
	defer resp.Body.Close()

	result := new(scienceGateAPIResponse)
	err = sonic.ConfigDefault.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("Error getting science gate events: unmarshalling error: %w", err)
	}

	events := make([]*model.Event, 0, len(result.Results))
	for _, r := range result.Results {
		events = append(events, s.mapToEvent(r))
	}
	return events, nil
}

// APIResponse представляет корневой ответ API
type scienceGateAPIResponse struct {
	Results []scienceGateResult `json:"results"`
}

// Result представляет одно событие из массива results
type scienceGateResult struct {
	Title       string   `json:"title"`
	Location    string   `json:"location"`
	Coordinates []string `json:"coordinates"`
	Periods     []period `json:"periods"`
	Site        string   `jspn:"organizerSite"`
}

// Period содержит временной диапазон события
type period struct {
	Lower string `json:"lower"`
	Upper string `json:"upper"`
}

func (s *ScienceGateService) mapToEvent(r scienceGateResult) *model.Event {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

	var startTime, endTime time.Time

	if len(r.Periods) > 0 {
		// Парсим первый период
		parsedStart, err := time.Parse(time.RFC3339Nano, r.Periods[0].Lower)
		if err == nil {
			startTime = parsedStart
		}

		parsedEnd, err := time.Parse(time.RFC3339Nano, r.Periods[0].Upper)
		if err == nil {
			endTime = parsedEnd
		}
	}

	// Если ивент начинается раньше текущего дня — ставим начало текущего дня
	if startTime.Before(todayStart) {
		startTime = todayStart
	}

	// Если ивент заканчивается после текущего дня — ставим конец текущего дня
	if endTime.After(todayEnd) {
		endTime = todayEnd
	}

	var coords [2]float64
	if len(r.Coordinates) >= 2 {
		if lat, err := strconv.ParseFloat(r.Coordinates[0], 64); err == nil {
			coords[0] = lat
		}
		if lon, err := strconv.ParseFloat(r.Coordinates[1], 64); err == nil {
			coords[1] = lon
		}
	} else {
		geoCodingResult, err := s.geocoder.GetCoords(r.Location)
		if err == nil {
			coords[0], coords[1] = geoCodingResult[0], geoCodingResult[1]
		}
	}

	return &model.Event{
		Title:        r.Title,
		StartTime:    startTime.Format("15:04"),
		EndTime:      endTime.Format("15:04"),
		Location:     r.Location,
		Coords:       coords,
		ExternalLink: r.Site,
		Tags:         []string{"НПВШ", "Наука", "Образование"},
	}
}
