package main

import (
	"context"
	"io"
	"net/http"
	"testing"
	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	//To send cancel signal, create ctx with cancel func
	ctx, cancel := context.WithCancel(context.Background())
	
	//Run http server
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	//Send http request and get response body to test http request and response works correctly
	input := "api/vocabularies"
	response, err := http.Get("http://localhost:8000/" + input)
	if err != nil {
		t.Errorf("Failed to get http response: %v", err)
	}
	defer response.Body.Close()
	got, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed to read http response body: %v", err)
	}

	//Compare response to expected value
	want := "Request Path: " + input
	if string(got) != want {
		t.Errorf("Want %q, but got %q", want, got)
	}

	//Send cancel signal to groutine which runs http server
	cancel()
	err = eg.Wait()
	if err != nil {
		t.Fatal(err)
	}
}

