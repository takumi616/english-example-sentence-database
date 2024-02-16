package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Compare json data
func AssertJSON(t *testing.T, wantJson, gotJson []byte) {
	t.Helper()

	var want, got any
	//Unmarshal expected json into go struct
	if err := json.Unmarshal(wantJson, &want); err != nil {
		t.Fatalf("Failed to unmarshal wantJson %q: %v", wantJson, err)
	}

	//Unmarshal real json into go struct
	if err := json.Unmarshal(gotJson, &got); err != nil {
		t.Fatalf("Failed to unmarshal gotJson %q: %v", gotJson, err)
	}

	//Compare
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Got differs: (-got +want)\n%s", diff)
	}
}

// Check if http response status code and response body
// have expected values.
func AssertResponse(t *testing.T, got *http.Response, wantStatus int, wantBody []byte) {
	t.Helper()
	t.Cleanup(func() { _ = got.Body.Close() })
	gotBody, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}

	//Check if status code is expected one.
	if got.StatusCode != wantStatus {
		t.Fatalf("Want status %d, but got %d, body: %q", wantStatus, got.StatusCode, gotBody)
	}

	if len(gotBody) == 0 && len(wantBody) == 0 {
		return
	}

	AssertJSON(t, wantBody, gotBody)
}

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read from %q: %v", path, err)
	}
	return bt
}
