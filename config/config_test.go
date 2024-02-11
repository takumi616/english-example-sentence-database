package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	want := 7000
	t.Setenv("APP_CONTAINER_PORT", fmt.Sprint(want))

	got, err := GetConfig() 
	if err != nil {
		t.Fatalf("Failed to get config: %v", err)
	}
	if got.Port != want {
		t.Errorf("Want %d, but got %d", want, got.Port)
	}
}