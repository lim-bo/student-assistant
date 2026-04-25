package shedule_service

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestGetSchedule(t *testing.T) {
	client := &EtuClient{
		http: &http.Client{Timeout: 10 * time.Second},
	}

	date := time.Now().AddDate(0, 0, 4)

	lessons, err := client.GetSchedule("3375", date)
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}

	fmt.Printf("Дата: %s\n", date.Format("2006-01-02"))
	fmt.Printf("Пар: %d\n\n", len(lessons))

	for _, l := range lessons {
		fmt.Printf("  [%s] %s — %s\n", l.Type, l.StartTime.Format("15:04"), l.EndTime.Format("15:04"))
		fmt.Printf("  %s\n", l.Title)

		if l.Location != nil {
			fmt.Printf("  %s\n", *l.Location)
		} else {
			fmt.Printf("  (дистант)\n")
		}

		if l.LocationLat != nil && l.LocationLon != nil {
			fmt.Printf("  coords: [%f, %f]\n", *l.LocationLat, *l.LocationLon)
		}

		fmt.Println()
	}
}
