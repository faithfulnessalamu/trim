package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/gorilla/mux"
)

var testDb = &memDatabase{data: make(map[string]string)}

func TestTrimHandlerExistentKey(t *testing.T) {
	// Test that a status ok is returned for a key that already exists
	// Create a request
	testPath := "/trim"
	payload := []byte(`{"url":"https://github.com"}`)

	req, err := http.NewRequest(http.MethodGet, testPath, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	actualURL := filepath.Join(baseURL, testPath)
	// Get hash
	trimHash := hash("https://github.com")
	// Save hash to test db
	testDb.save(trimHash, actualURL)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/trim", TrimHandler(testDb))
	router.ServeHTTP(rr, req)

	// Do validations
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code, got: %v, want: %v", status, http.StatusOK)
	}
}

func TestTrimHandlerBadURL(t *testing.T) {
	// Test that a bad request is returned for a bad url payload
	// Create a request
	testPath := "/trim"
	payload := []byte(`{"url":"https:git"}`)

	req, err := http.NewRequest(http.MethodGet, testPath, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/trim", TrimHandler(testDb))
	router.ServeHTTP(rr, req)

	// Do validations
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code, got: %v, want: %v", status, http.StatusBadRequest)
	}
}

func TestTrimHandler(t *testing.T) {
	// Test that a status OK is returned for a nonexistent key
	// Create a request
	testPath := "/trim"
	payload := []byte(`{"url":"https://github.com/thealamu/trim"}`)

	req, err := http.NewRequest(http.MethodGet, testPath, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/trim", TrimHandler(testDb))
	router.ServeHTTP(rr, req)

	// Do validations
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code, got: %v, want: %v", status, http.StatusOK)
	}
}

func TestRedirectHandler(t *testing.T) {
	// Test that the redirect handler actually redirects.
	// Use a mock database and insert the key and value beforehand
	// Create a request
	testPath := "/a1b2c3d4"
	req, err := http.NewRequest(http.MethodGet, testPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Insert data into the db
	if err := testDb.save("a1b2c3d4", "localhost:8080"); err != nil {
		t.Errorf("Could not save to mock database: %s", err.Error())
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a handler func using RedirectHandler and mux to capture url vars
	router := mux.NewRouter()
	router.HandleFunc("/{hash:[a-fA-F0-9]{8,8}}", RedirectHandler(testDb))

	// Record the response
	router.ServeHTTP(rr, req)

	// Do validations
	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code, got: %v, want: %v", status, http.StatusOK)
	}

	// Check body
	expected := "Redirecting now to localhost:8080"
	if got := rr.Body.String(); got != expected {
		t.Errorf("Unexpected body, got: %v, expected: %v", got, expected)
	}
}

func TestRedirectHandlerNotFound(t *testing.T) {
	// Test that the redirect handler
	//returns a 404 in the case that the key is not in the database
	// Create a request
	testPath := "/d4c3b2a1"
	req, err := http.NewRequest(http.MethodGet, testPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a handler func using RedirectHandler and mux to capture url vars
	router := mux.NewRouter()
	router.HandleFunc("/{hash:[a-fA-F0-9]{8,8}}", RedirectHandler(testDb))

	// Record the response
	router.ServeHTTP(rr, req)

	// Do validations
	// Check status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Wrong status code, got: %v, want: %v", status, http.StatusNotFound)
	}
}
