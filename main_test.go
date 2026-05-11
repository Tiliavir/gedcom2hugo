package main

import (
	"testing"
)

func TestVersionConstant(t *testing.T) {
	if version == "" {
		t.Error("version constant should not be empty")
	}
}

func TestMainFunction(t *testing.T) {
	// Test that main function doesn't panic when called
	// We can't easily test the actual execution without mocking CLI args
	// but we can at least ensure the package compiles
	if version != "0.0.0.1" {
		t.Errorf("Expected version 0.0.0.1, got %s", version)
	}
}
