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

	os.Setenv("ENV", "test")
	defer os.Unsetenv("ENV")

	os.Setenv("TW_ACC_SID", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	defer os.Unsetenv("TW_ACC_SID")

	os.Setenv("TW_AUTH_TOKEN", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	defer os.Unsetenv("TW_AUTH_TOKEN")

	os.Setenv("TW_PHONE_NUMBER", "+15005550006")
	defer os.Unsetenv("TW_PHONE_NUMBER")

	app.LoggerInit()
	cfg.LoadConfig()

	if cfg.port != "8080" {
		t.Errorf("want port %s; got %s", "8080", cfg.port)
	}

	if cfg.sender_email != "test@example.com" {
		t.Errorf("want %s get %s", "test@example.com", cfg.sender_email)
	}

	if cfg.env != "test" {
		t.Errorf("want %s get %s", "test", cfg.env)
	}

	if cfg.tw_acc_sid != "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
		t.Errorf("want %s get %s", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", cfg.tw_acc_sid)
	}

	if cfg.tw_auth_token != "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
		t.Errorf("want %s get %s", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", cfg.tw_auth_token)
	}

	if cfg.tw_phone_number != "+15005550006" {
		t.Errorf("want %s get %s", "+15005550006", cfg.tw_phone_number)
	}

}

func TestLoadConfigMissingEnv(t *testing.T) {
	os.Unsetenv("PORT")
	defer os.Setenv("PORT", "8080")
	os.Unsetenv("SENDER_MAIL")
	defer os.Setenv("SENDER_MAIL", "test@example.com")
	os.Unsetenv("ENV")
	defer os.Setenv("ENV", "test")
	os.Unsetenv("TW_ACC_SID")
	defer os.Setenv("TW_ACC_SID", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	os.Unsetenv("TW_AUTH_TOKEN")
	defer os.Setenv("TW_AUTH_TOKEN", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	os.Unsetenv("TW_PHONE_NUMBER")
	defer os.Setenv("TW_PHONE_NUMBER", "+15005550006")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected function to panic but it did not panic.")
		}
	}()

	cfg.LoadConfig()
}

func TestWriteJSON(t *testing.T) {

	data := envelope{
		"message": "email send successfully.",
	}

	rr := httptest.NewRecorder()

	headers := http.Header{}
	headers.Set("X-Custom-Header", "value")

	err := writeJSON(rr, http.StatusOK, data, headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{"message":"email send successfully."}` + "\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("want body %q; got %q", expectedBody, rr.Body.String())
	}

	expectedHeader := "value"
	if rr.Header().Get("X-Custom-Header") != expectedHeader {
		t.Errorf("want header %q; got %q", expectedHeader, rr.Header().Get("X-Custom-Header"))
	}
}
