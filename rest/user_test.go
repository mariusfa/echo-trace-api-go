package rest

import (
	"bytes"
	"echo/biz"
	"echo/biz/domain"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	userRepoFake := &biz.UserRepositoryFake{}
	router := SetupServicesControllers(userRepoFake)

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

	if len(userRepoFake.Users) != 1 {
		t.Errorf("User is not inserted")
	}
	if userRepoFake.Users[0].Name != "testuser" {
		t.Errorf("User is not inserted")
	}
	if userRepoFake.Users[0].HashedPassword == "testpass" {
		t.Errorf("Password is not hashed")
	}
	if userRepoFake.Users[0].ApiToken == "" {
		t.Errorf("ApiToken is not generated")
	}
}

func TestRegisterConflictUsername(t *testing.T) {
	user := domain.User{
		Name:           "testuser",
		HashedPassword: "testpass",
	}
	userRepoFake := &biz.UserRepositoryFake{}
	userRepoFake.Users = append(userRepoFake.Users, user)
	router := SetupServicesControllers(userRepoFake)

	// Create a request to send to the above route
	loginData := map[string]string{"username": "testuser", "password": "testpass2"}
	loginJSON, _ := json.Marshal(loginData)
	request := httptest.NewRequest("POST", "/user/register", bytes.NewBuffer(loginJSON))
	request.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response from the server
	response := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(response, request)

	// check length of users
	if len(userRepoFake.Users) == 2 {
		t.Errorf("Duplicate user inserted")
	}


	if response.Code != http.StatusConflict {
		t.Errorf("Response code is %v", response.Code)
	}
}

func TestValidate(t *testing.T) {
    // TODO implement
}
