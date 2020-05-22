package main

import (
	"testing"

	"github.com/aditya43/chess/chess"
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

// Test if error is returned when user inputs invalid piece type and invalid position
func TestForInvalidPieceTypeAndPositionInput(t *testing.T) {
	err := validateInput([]string{"adi", "x1"})

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

// Test if no error is returned for case-insensitive user input
func TestForCaseInsensitiveUserInput(t *testing.T) {
	inputs := [5][]string{
		{"qUeEn", "d4"},
		{"rOOk", "c8"},
		{"Bishop", "a1"},
		{"paWn", "h5"},
		{"kIng", "e7"},
	}

	for _, v := range inputs {
		err := validateInput([]string{v[0], v[1]})

		if err != nil {
			t.Fatalf("Failed to validate case-insensitive user input! Error: %v", err)
		}
	}
}

// Benchmark test for validateInput() func
func BenchmarkUserInputValidator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateInput([]string{"king", "a1"})
	}
}

// Benchmark test for CreateChessBoard() func
func BenchmarkCreateChessBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = chess.CreateChessBoard()
	}
}

// Benchmark test for generating move positions for a piece
func BenchmarkGenerateMovePositionsForPiece(b *testing.B) {
	cb := chess.CreateChessBoard()
	p := chess.CreatePiece(29, "pawn")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cb.PlacePiece(1, p)
	}
}

// Benchmark create chessboard and create piece
func BenchmarkCreateChessBoardAndPiece(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = chess.CreateChessBoard()
		_ = chess.CreatePiece(29, "pawn")
	}
}
