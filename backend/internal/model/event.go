package model

import (
	"github.com/bytedance/sonic"
)

type Event struct {
	Title        string     `json:"title"`
	StartTime    string     `json:"time_start"`
	EndTime      string     `json:"time_end"`
	Location     string     `json:"location"`
	Coords       [2]float64 `json:"coords"`
	Tags         []string   `json:"tags,omitempty"`
	ExternalLink string     `json:"external_link"`
}

func (e *Event) JsonString() string {
	str, _ := sonic.ConfigDefault.MarshalToString(e)
	return str
}
