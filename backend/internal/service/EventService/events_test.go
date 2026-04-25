package eventservice_test

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/student-assistant/internal/model"
	eventservice "github.com/student-assistant/internal/service/EventService"
	geocoding "github.com/student-assistant/pkg/GeoCoding"
)

func TestGetEvents(t *testing.T) {
	opts := eventservice.Options{
		Date: time.Now().Truncate(time.Hour * 24),
	}

	services := []eventservice.EventService{
		eventservice.NewKudago(),
		eventservice.NewScienceGate(geocoding.New()),
	}

	events := make([]*model.Event, 0)
	for _, s := range services {
		result, err := s.GetEvents(opts)
		assert.NoError(t, err)
		events = append(events, result...)
	}

	require.NotEqual(t, 0, len(events))
	for _, e := range events {
		log.Println(e.JsonString())
	}
}
