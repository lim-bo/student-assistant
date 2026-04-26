package eventservice

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/student-assistant/internal/config"
	"github.com/student-assistant/internal/model"
)

const (
	eventsOfTheDayEndpoint = "/events-of-the-day"
)

// Имплементация EventService под "Куда пойти"
type KudaGoService struct {
}

func NewKudago() *KudaGoService {
	return &KudaGoService{}
}

// Собираем ивенты
func (s *KudaGoService) GetEvents(opts Options) ([]*model.Event, error) {
	url := config.KudaGoAPIBaseUrl + eventsOfTheDayEndpoint
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error getting kudago events: %w", err)
	}

	q := req.URL.Query()
	q.Add("date", opts.Date.Format(time.DateOnly))
	q.Add("expand", "object,place")
	q.Add("location", string(opts.Loc))

	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error getting kudago events: %w", err)
	}
	defer resp.Body.Close()

	result := new(kudagoAPIResponse)
	err = sonic.ConfigDefault.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("Error getting kudago events: unmarshalling error: %w", err)
	}

	events := make([]*model.Event, 0, len(result.Results))

	for _, r := range result.Results {
		events = append(events, s.mapToEvent(r, opts.Date))
	}

	return events, nil
}

// APIResponse представляет корневой ответ API
type kudagoAPIResponse struct {
	Results []kudagoResult `json:"results"`
}

// Result представляет одно событие из массива results
type kudagoResult struct {
	Object object `json:"object"`
}

// Object содержит основную информацию о событии
type object struct {
	Title     string    `json:"title"`
	ItemURL   string    `json:"item_url"`
	Place     place     `json:"place"`
	Daterange dateRange `json:"daterange"`
}

// Place содержит информацию о месте проведения
type place struct {
	Address string `json:"address"`
	Coords  coords `json:"coords"`
}

// Coords представляет географические координаты
type coords struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// DateRange содержит информацию о временном диапазоне события
type dateRange struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

// Маппинг
func (s *KudaGoService) mapToEvent(r kudagoResult, day time.Time) *model.Event {
	var coords [2]float64
	if r.Object.Place.Coords.Lat != 0 || r.Object.Place.Coords.Lon != 0 {
		coords = [2]float64{r.Object.Place.Coords.Lat, r.Object.Place.Coords.Lon}
	}

	now := day
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

	startTime := time.Unix(r.Object.Daterange.Start, 0)
	endTime := time.Unix(r.Object.Daterange.End, 0)

	// Если ивент начинается раньше текущего дня - ставим начало текущего дня
	if startTime.Before(todayStart) {
		startTime = todayStart
	}

	// Если ивент заканчивается после текущего дня - ставим конец текущего дня
	if endTime.After(todayEnd) {
		endTime = todayEnd
	}

	return &model.Event{
		Title:        r.Object.Title,
		StartTime:    startTime,
		EndTime:      endTime,
		Location:     r.Object.Place.Address,
		Coords:       coords,
		ExternalLink: r.Object.ItemURL,
		Tags:         []string{"KudaGo", "Досуг"},
	}
}
