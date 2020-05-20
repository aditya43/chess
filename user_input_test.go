package main

import (
	"testing"
)

// Test if error is returned when user inputs a single word
func TestForSingleWordInput(t *testing.T) {
	err := validateInput([]string{"singleword"})

	if err == nil {
		t.Fatalf("Failed to validate single word user input! Error: %v", err)
	}
}
