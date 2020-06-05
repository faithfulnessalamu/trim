package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var testDb = &memDatabase{data: make(map[string]string)}

func TestRedirectHandler(t *testing.T) {
	// Create a request
	testPath := "/a1b2c3d4"
	req, err := http.NewRequest(http.MethodGet, testPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a handler func using RedirectHandler and mux to capture url vars
	router := mux.NewRouter()
	router.HandleFunc("/{hash:[a-zA-Z0-9]{8,8}}", RedirectHandler(testDb))

	// Record the response
	router.ServeHTTP(rr, req)

	// Do validations
	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code, got: %v, want: %v", status, http.StatusOK)
	}

	// Check body
	expected := "Redirecting now from a1b2c3d4"
	if got := rr.Body.String(); got != expected {
		t.Errorf("Unexpected body, got: %v, expected: %v", got, expected)
	}
}
