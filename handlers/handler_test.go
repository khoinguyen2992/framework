package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRenderJSON(t *testing.T) {
	req, err := http.NewRequest("GET", "/render", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := func(rw http.ResponseWriter, req *http.Request) {
		RenderJSON(rw, http.StatusOK, "ok")
	}

	handler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected: %d, got: %d", http.StatusOK, recorder.Code)
	}

	expectedData, _ := json.Marshal(map[string]interface{}{
		"data": "ok",
	})

	if string(recorder.Body.Bytes()) != string(expectedData) {
		t.Errorf("Expected: %s, got: %s", string(expectedData), string(recorder.Body.Bytes()))
	}
}

func TestRenderError(t *testing.T) {
	req, err := http.NewRequest("GET", "/render", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := func(rw http.ResponseWriter, req *http.Request) {
		RenderError(rw, http.StatusBadRequest, "error")
	}

	handler(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected: %d, got: %d", http.StatusOK, recorder.Code)
	}

	expectedData, _ := json.Marshal(map[string]string{
		"error": "error",
	})

	if string(recorder.Body.Bytes()) != string(expectedData) {
		t.Errorf("Expected: %s, got: %s", string(expectedData), string(recorder.Body.Bytes()))
	}
}
