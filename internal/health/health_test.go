package health

import (
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHealthHandler(t *testing.T) {
	r := chi.NewRouter()
	RegisterRoutes(r)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	if contentType := rec.Header().Get("Content-Type"); !strings.HasPrefix(contentType, "application/json") {
		t.Errorf("expected content type 'application/json', got '%s'", contentType)
	}

	var healthResp HealthResponse
	err := json.NewDecoder(rec.Body).Decode(&healthResp)
	if err != nil {
		t.Fatalf("expected valid JSON response, got error: %v", err)
	}
	if healthResp.Status != "ok" {
		t.Errorf("expected status 'ok', got '%s'", healthResp.Status)
	}
}
