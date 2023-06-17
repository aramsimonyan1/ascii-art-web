package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the homeHandler function directly, passing in the ResponseRecorder and Request
	homeHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", rr.Code)
	}

	// Add more assertions as needed
}

func TestAsciiArtHandler(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("POST", "/ascii-art", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the asciiArtHandler function directly, passing in the ResponseRecorder and Request
	asciiArtHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", rr.Code)
	}

	// Add more assertions as needed
}

func TestGenerateASCIIArt(t *testing.T) {
	// Create test inputs
	text := "Hello"
	banner := "standard"

	// Call the generateASCIIArt function with the test inputs
	result, err := generateASCIIArt(text, banner)
	// Check for errors
	if err != nil {
		t.Errorf("Failed to generate ASCII art: %v", err)
	}

	// Check if the result is not empty
	if result == "" {
		t.Errorf("Expected non-empty result, but got an empty string")
	}
}
