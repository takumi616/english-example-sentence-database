package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	t.Skip("Refactoring now")

	//To send cancel signal, create ctx with cancel func
	ctx, cancel := context.WithCancel(context.Background())

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to get a listener with port: %v", err)
	}

	//Run http server
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	//Send http request and get response body 
	//to test that http request and response works correctly
	input := "api/vocabularies"
	url := fmt.Sprintf("http://%s/%s", listener.Addr().String(), input)
	t.Logf("Http request URL: %q", url)
	response, err := http.Get(url)
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

