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

// Test if no error is returned for a correct input
func TestForCorrectUserInput(t *testing.T) {
	inputs := [5][]string{
		{"queen", "d4"},
		{"rook", "c8"},
		{"bishop", "a1"},
		{"pawn", "h5"},
		{"king", "e7"},
	}

	for _, v := range inputs {
		err := validateInput([]string{v[0], v[1]})

		if err != nil {
			t.Fatalf("Failed to validate correct user input! Error: %v", err)
		}
	}
}
