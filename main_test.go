package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	app := &application{}

	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "api/v1/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.healthCheck(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rr.Code)
	}

}
