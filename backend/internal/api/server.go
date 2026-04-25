package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	eventservice "github.com/student-assistant/internal/service/event_service"
	shedule_service "github.com/student-assistant/internal/service/schedule_service"
	geocoding "github.com/student-assistant/pkg/GeoCoding"
)

type Server struct {
	mux *chi.Mux
}

func New() *Server {
	return &Server{
		mux: chi.NewMux(),
	}
}

func (s *Server) mountEndpoints() {
	s.mux.Route("/api", func(r chi.Router) {
		r.Use(s.CORSMiddleware)

		h := NewHandler(
			shedule_service.NewEtuClient(),
			[]eventservice.EventService{
				eventservice.NewKudago(),
				eventservice.NewScienceGate(geocoding.New()),
			},
		)
		r.Get("/schedule", h.Plan)
	})
}

func (s *Server) Run(address string, port int) error {
	s.mountEndpoints()

	return http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), s.mux)
}
