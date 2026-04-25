package shedule_service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/student-assistant/internal/config"
	"github.com/student-assistant/internal/model"
)

type EtuClient struct {
	http *http.Client
}

func NewEtuClient() *EtuClient {
	return &EtuClient{
		http: &http.Client{Timeout: 10 * time.Second},
	}
}

type etuScheduleResponse map[string]etuScheduleEntry

type etuScheduleEntry struct {
	Group string            `json:"group"`
	Days  map[string]etuDay `json:"days"`
}

type etuDay struct {
	Name    string      `json:"name"`
	Lessons []etuLesson `json:"lessons"`
}

type etuLesson struct {
	Name      string `json:"name"`
	Week      string `json:"week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Room      string `json:"room"`
	Form      string `json:"form"`
}

func (c *EtuClient) GetSchedule(groupNumber string, date time.Time) ([]model.ScheduleLesson, error) {
	url := fmt.Sprintf("%s/mobile/schedule?groupNumber=%s", config.EtuBase, groupNumber)
	body, err := c.get(url)
	if err != nil {
		return nil, fmt.Errorf("GetSchedule request: %w", err)
	}

	var resp etuScheduleResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("GetSchedule parse: %w", err)
	}

	entry, ok := resp[groupNumber]
	if !ok {
		return nil, fmt.Errorf("group %q not found", groupNumber)
	}

	etuDayIndex := weekdayToEtu(date.Weekday())
	day, ok := entry.Days[etuDayIndex]
	if !ok || len(day.Lessons) == 0 {
		return []model.ScheduleLesson{}, nil
	}

	weekNumber := getWeekNumber(date)
	return convertLessons(day.Lessons, date, weekNumber), nil
}

func convertLessons(raw []etuLesson, date time.Time, currentWeek int) []model.ScheduleLesson {
	y, m, d := date.Date()
	loc := date.Location()
	result := make([]model.ScheduleLesson, 0)

	seen := make(map[string]bool)

	for _, l := range raw {
		if l.Week != fmt.Sprintf("%d", currentWeek) {
			continue
		}

		key := l.Name + l.StartTime
		if seen[key] {
			continue
		}
		seen[key] = true
		formOfStudy := formOfStudy(l.Form)

		var location *string
		var lat, lon *float64

		if l.Name == "Военно-морская подготовка" { // захардкодил
			addr := "Улица Профессора Попова, 37Б лит А"
			la, lo := 59.971582, 30.296756
			location = &addr
			lat = &la
			lon = &lo
		} else if formOfStudy != "remote" {
			addr := fmt.Sprintf("%s, ауд. %s", roomAddress(l.Room), l.Room)
			la, lo := roomCoords(l.Room)
			location = &addr
			lat = &la
			lon = &lo
		}

		result = append(result, model.ScheduleLesson{
			Title:       l.Name,
			Location:    location, // null в JSON если remote
			Type:        formOfStudy,
			StartTime:   parseHHMM(l.StartTime, y, m, d, loc),
			EndTime:     parseHHMM(l.EndTime, y, m, d, loc),
			LocationLat: lat,
			LocationLon: lon,
		})
	}

	return result
}

func roomCoords(room string) (lat, lon float64) {
	if len(room) == 0 {
		return 0, 0
	}
	switch room[0] {
	case '1':
		return 59.971614, 30.320171
	case '2':
		return 59.971380, 30.321554
	case '3':
		return 59.972814, 30.324519
	case '4':
		return 59.973709, 30.323502
	case '5':
		return 59.971947, 30.324303
	case '6':
		return 59.971582, 30.296756
	}
	return 0, 0
}

func roomAddress(room string) string {
	if len(room) == 0 {
		return ""
	}
	switch room[0] {
	case '1':
		return "улица Профессора Попова, 5Б"
	case '2':
		return "улица Профессора Попова, 5к2"
	case '3':
		return "Инструментальная улица, 2"
	case '4':
		return "Инструментальная улица, 4"
	case '5':
		return "улица Профессора Попова, 5к5"
	case '6':
		return "Улица Профессора Попова, 37Б лит А"
	}
	return ""
}

func weekdayToEtu(wd time.Weekday) string {
	m := map[time.Weekday]string{
		time.Monday:    "0",
		time.Tuesday:   "1",
		time.Wednesday: "2",
		time.Thursday:  "3",
		time.Friday:    "4",
		time.Saturday:  "5",
		time.Sunday:    "6",
	}
	return m[wd]
}

func formOfStudy(form string) string {
	switch form {
	case "standard":
		return "class"
	case "distant", "online":
		return "remote"
	}
	return ""
}

func getWeekNumber(date time.Time) int {
	year := date.Year()
	if date.Month() < time.September {
		year--
	}
	sept := time.Date(year, time.September, 1, 0, 0, 0, 0, date.Location())
	for sept.Weekday() != time.Monday {
		sept = sept.AddDate(0, 0, 1)
	}

	days := int(date.Sub(sept).Hours() / 24)
	week := (days / 7) + 1
	if week%2 == 0 {
		return 2
	}
	return 1
}

func parseHHMM(hhmm string, y int, m time.Month, d int, loc *time.Location) time.Time {
	var hour, min int
	fmt.Sscanf(hhmm, "%d:%d", &hour, &min)
	return time.Date(y, m, d, hour, min, 0, 0, loc)
}

func (c *EtuClient) get(url string) ([]byte, error) {
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: status %d", url, resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
