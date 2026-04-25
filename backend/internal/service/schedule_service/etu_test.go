package shedule_service

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestGetDaySchedule(t *testing.T) {
	client := &EtuClient{
		http: &http.Client{Timeout: 10 * time.Second},
	}

	date := time.Date(2026, 4, 25, 0, 0, 0, 0, time.Local)

	schedule, err := client.GetDaySchedule("3376", date)
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}

	fmt.Printf("Дата: %s\n", schedule.Date.Format("2006-01-02"))
	fmt.Printf("Пар: %d\n\n", len(schedule.Lessons))

	for _, l := range schedule.Lessons {
		fmt.Printf("  [%s] %s — %s\n", l.Type, l.StartTime.Format("15:04"), l.EndTime.Format("15:04"))
		fmt.Printf("  %s\n", l.Title)
		fmt.Printf("  %s\n", l.Location)
		fmt.Printf("  coords: [%f, %f]\n\n", l.LocationLat, l.LocationLon)
	}
}
