package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

/*
Function checks whether your project handles the HTTP status code 404 - Not Found correctly.
It simulates a request to a non-existent path (endpoint) and check if the application returns a "404 Not Found" status code.
*/
func TestNotFoundHandler(t *testing.T) {
	// We create a GET request to the /non-existing-endpoint path, which is a non-existing endpoint in your application.
	req, err := http.NewRequest("GET", "/non-existent", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	/* A method call that directly invokes the ServeHTTP method of the default HTTP request multiplexer (http.DefaultServeMux).
	This approach bypasses the registered handler functions and directly uses the default multiplexer to handle the request.
	It can be useful in certain scenarios, such as when you want to simulate requests to specific paths that are not handled
	by your registered handlers.
	Since there is no registered handler for the requested path, the default multiplexer will return a "404 Not Found" response.
	*/
	http.DefaultServeMux.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, but got %d", rr.Code)
	}
}

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
}

/* We create the request using http.NewRequest("POST", "/ascii-art", 1), but we set
the "Content-Type" header to "application/x-www-form-urlencoded" explicitly.
This indicates that the request will contain form data. */
func TestBadRequestHandler(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("POST", "/ascii-art", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header to indicate form data
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the asciiArtHandler function directly, passing in the ResponseRecorder and Request
	asciiArtHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, but got %d", rr.Code)
	}
}

/*
Function is responsible for testing the handling of an internal server error (HTTP status code 500) by the server.
*/
func TestInternalServerErrorHandler(t *testing.T) {
	// We create a ResponseRecorder (rr) to capture the response.
	rr := httptest.NewRecorder()

	// Then, we simulate an internal server error by calling http.Error and passing the rr recorder,
	// the error message ("Internal Server Error"), and the status code http.StatusInternalServerError.
	http.Error(rr, "Internal Server Error", http.StatusInternalServerError)

	// We check the response status code (rr.Code) and compare it with the expected value (http.StatusInternalServerError).
	// If the status code doesn't match the expected value, we report an error indicating the mismatch.
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, but got %d", rr.Code)
	}
}

/* Does the test send a valid request to the asciiArtHandler function?
We create a formData variable that holds the required form data with "inputText" and "banners" fields.
We then encode the form data and set it as the request body using strings.NewReader(formData.Encode()).
Additionally, we set the "Content-Type" header to "application/x-www-form-urlencoded" to indicate the form data format.
*/
func TestAsciiArtHandler(t *testing.T) {
	// Create a request to pass to the handler
	formData := url.Values{
		"inputText": []string{"example text"},
		"banners":   []string{"shadow"},
	}
	req, err := http.NewRequest("POST", "/ascii-art", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the asciiArtHandler function directly, passing in the ResponseRecorder and Request
	asciiArtHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", rr.Code)
	}
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
