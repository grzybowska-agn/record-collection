package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireTokenRejectsMissingToken(t *testing.T) {
	app := application{apiToken: "test-secret"}
	called := false

	handler := app.requireToken(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	request := httptest.NewRequest(http.MethodPost, "/records", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", response.Code)
	}

	if called {
		t.Error("protected handler ran without authentication")
	}
}

func TestRequireTokenAcceptsValidToken(t *testing.T) {
	app := application{apiToken: "test-secret"}

	handler := app.requireToken(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	request := httptest.NewRequest(http.MethodPost, "/record", nil)
	request.Header.Set("Authorization", "Bearer test-secret")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", response.Code)
	}
}
