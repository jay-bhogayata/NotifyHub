package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var cfg = &config{}
var app = &application{
	config: *cfg,
}

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

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

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "8080")
	defer os.Unsetenv("PORT")

	os.Setenv("SENDER_EMAIL", "test@example.com")
	defer os.Unsetenv("SENDER_EMAIL")

	app.LoggerInit()
	cfg.LoadConfig()

	if cfg.port != "8080" {
		t.Errorf("want port %s; got %s", "8080", cfg.port)
	}

	if cfg.sender_mail != "test@example.com" {
		t.Errorf("want %s get %s", "test@example.com", cfg.sender_mail)
	}
}

func TestLoadConfigMissingEnv(t *testing.T) {
	os.Unsetenv("PORT")
	defer os.Setenv("PORT", "8080")
	os.Unsetenv("SENDER_MAIL")
	defer os.Setenv("SENDER_MAIL", "test@example.com")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected function to panic but it did not panic.")
		}
	}()

	cfg.LoadConfig()
}
