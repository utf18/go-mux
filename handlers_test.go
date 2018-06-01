package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestYourHandler(t *testing.T) {
	var r *http.Request
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	w := httptest.NewRecorder()

	YourHandler(w, r)
	actualResponseBody := w.Body.String()
	expectedResponseBody := "this is not the endpoint you are looking for... try /metrics"
	if actualResponseBody != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actualResponseBody, expectedResponseBody)
	}
}
