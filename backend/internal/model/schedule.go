package model

import "time"

type ScheduleLesson struct {
	Title       string    `json:"title"`
	Location    *string   `json:"location,omitempty"`
	Type        string    `json:"type"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	LocationLat *float64  `json:"location_lat,omitempty"`
	LocationLon *float64  `json:"location_lon,omitempty"`
}
