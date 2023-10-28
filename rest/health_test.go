package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	HealthController := HealthController{}
	server := SetupHandlers(HealthController)

	// Create a request to send to the above route
	request := httptest.NewRequest("GET", "/health", nil)
	// Create a response recorder to record the response from the server
	response := httptest.NewRecorder()
	// Perform the request
	server.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}
}
