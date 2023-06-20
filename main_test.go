package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/non-existent", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, but got %d", rr.Code)
	}
}

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	homeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", rr.Code)
	}
}

func TestBadRequestHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/ascii-art", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	asciiArtHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, but got %d", rr.Code)
	}
}

func TestInternalServerErrorHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	http.Error(rr, "Internal Server Error", http.StatusInternalServerError)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, but got %d", rr.Code)
	}
}

func TestAsciiArtHandler(t *testing.T) {
	formData := url.Values{
		"inputText": []string{"example text"},
		"banners":   []string{"shadow"},
	}
	req, err := http.NewRequest("POST", "/ascii-art", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	asciiArtHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", rr.Code)
	}
}

func TestGenerateASCIIArt(t *testing.T) {
	text := "Hello"
	banner := "standard"

	result, err := generateASCIIArt(text, banner)
	if err != nil {
		t.Errorf("Failed to generate ASCII art: %v", err)
	}

	if result == "" {
		t.Errorf("Expected non-empty result, but got an empty string")
	}
}
