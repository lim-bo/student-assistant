package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/student-assistant/internal/api"
	eventservice "github.com/student-assistant/internal/service/event_service"
	shedule_service "github.com/student-assistant/internal/service/schedule_service"
	geocoding "github.com/student-assistant/pkg/GeoCoding"
)

func TestHandlerPlan(t *testing.T) {
	h := api.NewHandler(
		shedule_service.NewEtuClient(),
		[]eventservice.EventService{
			eventservice.NewKudago(),
			eventservice.NewScienceGate(geocoding.New()),
		},
	)

	req := httptest.NewRequest(http.MethodGet, "/api/schedule", nil)

	q := req.URL.Query()
	q.Add("group", "3376")
	q.Add("date", "2026-04-27")
	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()

	h.Plan(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидали 200, получили %d: %s", w.Code, w.Body.String())
	}

	var resp api.PlanResponse
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
	h := api.NewHandler(shedule_service.NewEtuClient(), nil)

	req := httptest.NewRequest(http.MethodGet, "/api/schedule", nil)

	q := req.URL.Query()
	q.Add("group", "3376")
	q.Add("date", "28-04-2026")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()

	h.Plan(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("ожидали 400, получили %d", w.Code)
	}
}

func TestHandlerPlan_BadGroup(t *testing.T) {
	h := api.NewHandler(shedule_service.NewEtuClient(), nil)

	req := httptest.NewRequest(http.MethodGet, "/api/schedule", nil)

	q := req.URL.Query()
	q.Add("date", "28-04-2026")
	w := httptest.NewRecorder()

	h.Plan(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("ожидали 400, получили %d", w.Code)
	}
}
