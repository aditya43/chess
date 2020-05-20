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

// Test if error is returned when user inputs an invalid piece name
func TestForInvalidPieceNameInput(t *testing.T) {
	err := validateInput([]string{"invalid-piece-name", "a1"})

	if err == nil {
		t.Fatalf("Failed to validate invalid piece name user input! Error: %v", err)
	}
}

// Test if error is returned when user inputs an invalid position
func TestForInvalidPositionInput(t *testing.T) {
	err := validateInput([]string{"rook", "x1"})

	if err == nil {
		t.Fatalf("Failed to validate invalid position user input! Error: %v", err)
	}
}
