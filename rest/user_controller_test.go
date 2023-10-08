package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	userController := &UserController{}
	healthController := &HealthController{}
	router := SetupRouter(healthController, userController)

	// Create a request to send to the above route
	loginData := map[string]string{"username": "testuser", "password": "testpass"}
	loginJSON, _ := json.Marshal(loginData)
	request := httptest.NewRequest("POST", "/user/register", bytes.NewBuffer(loginJSON))
	request.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response from the server
	response := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}
}