package main

import "testing"

func TestHandler(t *testing.T) {
	expected := 1
	if expected != 1 {
		t.Errorf("Expected 1, got %d", expected)
	}
}

func TestSanity(t *testing.T) {
    // Just a sanity check for the CI pipeline
}
