package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateEvent_MissingSource_ReturnsBadRequest(t *testing.T) {

	body := `{
		"event_id": "evt-001",
		"event_type": "order.delivered",
		"payload": {"order_id": "order-123"}
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/events",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	CreateEvent(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d; body: %s",
			http.StatusBadRequest,
			rec.Code,
			rec.Body.String(),
		)
	}
}

func TestCreateEvent_ArrayPayload_ReturnsBadRequest(t *testing.T) {
	body := `{
		"source": "shipping-service",
		"event_id": "evt-002",
		"event_type": "order.delivered",
		"payload": ["order-123"]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/events",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	CreateEvent(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d; body: %s",
			http.StatusBadRequest,
			rec.Code,
			rec.Body.String(),
		)
	}
}

func TestCreateEvent_ValidEvent_ReturnsCreated(t *testing.T) {
	body := `{
		"source": "shipping-service",
		"event_id": "evt-003",
		"event_type": "order.delivered",
		"payload": {"order_id": "order-123"}
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/events",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	CreateEvent(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d; body: %s",
			http.StatusCreated,
			rec.Code,
			rec.Body.String(),
		)
	}
}
