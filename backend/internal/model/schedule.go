package model

import "time"

type ScheduleLesson struct {
	Title       string    `json:"title"`
	Location    string    `json:"location"`
	Type        string    `json:"type"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	LocationLat float64   `json:"location_lat"`
	LocationLon float64   `json:"location_lon"`
}

type DaySchedule struct {
	Date    time.Time        `json:"date"`
	Lessons []ScheduleLesson `json:"lessons"`
}
