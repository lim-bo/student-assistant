package model

type TimelineItem struct {
	TimeStart    string      `json:"time_start"`
	TimeEnd      string      `json:"time_end"`
	Type         string      `json:"type"`
	Title        string      `json:"title"`
	Location     *string     `json:"location,omitempty"`
	Coords       *[2]float64 `json:"coords,omitempty"`
	ExternalLink string      `json:"external_link,omitempty"`
	Tags         []string    `json:"tags,omitempty"`
}
