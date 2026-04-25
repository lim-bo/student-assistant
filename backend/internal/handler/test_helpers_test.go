package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	eventservice "github.com/student-assistant/internal/service/event_service"
	shedule_service "github.com/student-assistant/internal/service/schedule_service"
	geocoding "github.com/student-assistant/pkg/GeoCoding"
)

func TestHandlerPlan(t *testing.T) {
	geocoder := geocoding.New()

	h := NewHandler(
		shedule_service.NewEtuClient(),
		[]eventservice.EventService{
			eventservice.NewKudago(),
			eventservice.NewScienceGate(geocoder),
		},
	)

	body, _ := json.Marshal(PlanRequest{
		Group: "3376",
		Date:  "2026-04-27",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/plan", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.HandlerPlan(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидали 200, получили %d: %s", w.Code, w.Body.String())
	}

	var resp PlanResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("не удалось распарсить ответ: %v", err)
	}

	t.Logf("Всего элементов: %d\n", len(resp.Timeline))
	for _, item := range resp.Timeline {
		data, _ := json.MarshalIndent(item, "  ", "  ")
		t.Logf("%s\n", data)
	}
}

func TestHandlerPlan_BadDate(t *testing.T) {
	h := NewHandler(shedule_service.NewEtuClient(), nil)

	body, _ := json.Marshal(PlanRequest{
		Group: "3376",
		Date:  "28-04-2026", // неправильный формат
	})

	req := httptest.NewRequest(http.MethodPost, "/api/plan", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.HandlerPlan(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("ожидали 400, получили %d", w.Code)
	}
}

func TestHandlerPlan_BadBody(t *testing.T) {
	h := NewHandler(shedule_service.NewEtuClient(), nil)

	req := httptest.NewRequest(http.MethodPost, "/api/plan", bytes.NewReader([]byte("not json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.HandlerPlan(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("ожидали 400, получили %d", w.Code)
	}
}
