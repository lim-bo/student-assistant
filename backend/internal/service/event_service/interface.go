package eventservice

import (
	"time"

	"github.com/student-assistant/internal/model"
)

type Location string

const (
	Spb    Location = "spb"
	Moscow Location = "msk"
)

type Options struct {
	Date time.Time
	Loc  Location
}

type EventService interface {
	GetEvents(opts Options) ([]*model.Event, error)
}
