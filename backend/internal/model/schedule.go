package model

import (
	"time"
)

type ScheduleLesson struct {
	Title       string    `json:"title"`
	Location    *string   `json:"location,omitempty"`
	Type        string    `json:"type"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	LocationLat *float64  `json:"location_lat,omitempty"`
	LocationLon *float64  `json:"location_lon,omitempty"`
}

func (l *ScheduleLesson) ToTimelineItem() *TimelineItem {
	return &TimelineItem{
		TimeStart: l.StartTime.Format("15:04"),
		TimeEnd:   l.EndTime.Format("15:04"),
		Type:      l.Type,
		Title:     l.Title,
		Location:  l.Location,
		Coords:    coordsFromLesson(*l),
	}
}

func coordsFromLesson(l ScheduleLesson) *[2]float64 {
	if l.LocationLat == nil || l.LocationLon == nil {
		return nil
	}
	coords := [2]float64{*l.LocationLat, *l.LocationLon}
	return &coords
}
