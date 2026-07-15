package main

import "testing"

func TestServerAddressUsesAddressEnvironmentVariable(t *testing.T) {
	t.Setenv("ADDRESS", "127.0.0.1:8081")

	if got := serverAddress(); got != "127.0.0.1:8081" {
		t.Fatalf("serverAddress() = %q, want %q", got, "127.0.0.1:8081")
	}
}
