package main

import (
	"net"
	"testing"
)

func TestRun(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to get a listener with port: %v", err)
	}
	t.Log(listener)
}
