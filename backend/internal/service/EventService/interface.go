package eventservice

import (
	"time"

	"github.com/student-assistant/internal/model"
)

type Options struct {
	Date time.Time
}

type EventService interface {
	GetEvents(opts Options) ([]*model.Event, error)
}
