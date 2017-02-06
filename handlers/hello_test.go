package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/render", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler := HelloHandler{}
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected: %d, got: %d", http.StatusOK, recorder.Code)
	}

	expectedData, _ := json.Marshal(map[string]interface{}{
		"data": map[string]string{
			"msg": "hello",
		},
	})

	if string(recorder.Body.Bytes()) != string(expectedData) {
		t.Errorf("Expected: %s, got: %s", string(expectedData), string(recorder.Body.Bytes()))
	}
}
