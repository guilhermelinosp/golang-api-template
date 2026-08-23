package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/guilhermelinosp/hellnet-lib-telemetry/telemetry"
)

func TestHealthHandler(t *testing.T) {
	ops, err := telemetry.New(telemetry.Options{ServiceName: "test", Enabled: false})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	ops.Health().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var hs telemetry.HealthStatus
	if err := json.Unmarshal(rec.Body.Bytes(), &hs); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if hs.Status != "ok" {
		t.Fatalf("expected status ok, got %q", hs.Status)
	}
}
