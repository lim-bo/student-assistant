package model

import (
	"time"

	"github.com/bytedance/sonic"
)

type Event struct {
	Title        string     `json:"title"`
	StartTime    time.Time  `json:"time_start"`
	EndTime      time.Time  `json:"time_end"`
	Location     string     `json:"location"`
	Coords       [2]float64 `json:"coords"`
	Tags         []string   `json:"tags,omitempty"`
	ExternalLink string     `json:"external_link"`
}

func (e *Event) JsonString() string {
	str, _ := sonic.ConfigDefault.MarshalToString(e)
	return str
}

func (e *Event) ToTimelineItem() *TimelineItem {
	var timelineCoords *[2]float64
	if e.Coords[0]+e.Coords[1] > 1 {
		timelineCoords = &e.Coords
	}

	return &TimelineItem{
		TimeStart:    e.StartTime.Format("15:04"),
		TimeEnd:      e.EndTime.Format("15:04"),
		Type:         "event",
		Title:        e.Title,
		Location:     &e.Location,
		Coords:       timelineCoords,
		ExternalLink: e.ExternalLink,
		Tags:         e.Tags,
	}
}
