package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetUpRouting(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	mux := SetUpRouting()
	//ServeHTTP dispatches the request to the handler
	//whose pattern most closely matches the request URL.
	mux.ServeHTTP(w, req)

	res := w.Result()
	//Close response body when this test function finishes.
	t.Cleanup(
		func() { _ = res.Body.Close() },
	)

	if res.StatusCode != http.StatusOK {
		t.Error("Want status code 200 but got ", res.StatusCode)
	}
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to response body: %v", err)
	}

	want := `{"Status": "ok"}`
	if string(got) != want {
		t.Errorf("Want %q, but got %q", want, got)
	}
}
