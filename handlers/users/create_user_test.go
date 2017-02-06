package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"framework/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockSuccessUserRepository struct{}

func (repo *mockSuccessUserRepository) Create(user *models.User) error {
	user.Id = 1
	return nil
}

type mockFailUserRepository struct{}

func (repo *mockFailUserRepository) Create(user *models.User) error {
	return errors.New("Cannot create user")
}

func TestCreateUserHandler(t *testing.T) {
	header := make(http.Header)
	header.Set("Content-Type", "application/json")
	data := `
   {
     "email":"test@example.com"
   }`
	req, _ := http.NewRequest("POST", "/users/", bytes.NewBufferString(data))
	req.Header = header

	recorder := httptest.NewRecorder()

	handler := CreateUserHandler{&mockSuccessUserRepository{}}
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected: %d, got: %d", http.StatusOK, recorder.Code)
	}

	body := map[string]models.User{}
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		t.Errorf("It should NOT get error: %s", err.Error())
	}

	user := body["data"]
	if user.Id != 1 {
		t.Errorf("Expected: 1, got: %d", user.Id)
	}

	if user.Email != "test@example.com" {
		t.Errorf("Expected: test@example.com, got: %s", user.Email)
	}
}

func TestCreateUserWhenRequestIsInvalid(t *testing.T) {
	header := make(http.Header)
	header.Set("Content-Type", "application/json")
	data := `
   {
     "email":""
   }`
	req, _ := http.NewRequest("POST", "/users/", bytes.NewBufferString(data))
	req.Header = header

	recorder := httptest.NewRecorder()

	handler := CreateUserHandler{&mockSuccessUserRepository{}}
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected: %d, got: %d", http.StatusOK, recorder.Code)
	}

	expectedData, _ := json.Marshal(map[string]string{
		"error": "email required, email invalid",
	})

	if string(recorder.Body.Bytes()) != string(expectedData) {
		t.Errorf("Expected: %s, got: %s", string(expectedData), string(recorder.Body.Bytes()))
	}
}

func TestCreateUserWhenDatabaseHasProblem(t *testing.T) {
	header := make(http.Header)
	header.Set("Content-Type", "application/json")
	data := `
   {
     "email":"test@example.com"
   }`
	req, _ := http.NewRequest("POST", "/users/", bytes.NewBufferString(data))
	req.Header = header

	recorder := httptest.NewRecorder()

	handler := CreateUserHandler{&mockFailUserRepository{}}
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected: %d, got: %d", http.StatusOK, recorder.Code)
	}

	expectedData, _ := json.Marshal(map[string]string{
		"error": "Cannot create user",
	})

	if string(recorder.Body.Bytes()) != string(expectedData) {
		t.Errorf("Expected: %s, got: %s", string(expectedData), string(recorder.Body.Bytes()))
	}
}
