package main

import (
	"reflect"
	"testing"
)

// Test if a chessboard has 64 cells in it
func TestChessBoardHas64Cells(t *testing.T) {
	if len(cb.cells) != 64 {
		t.Fatalf("Failed to test if a chessboard has 64 cells in it, Number of cells: %v", len(cb.cells))
	}
}

// Test if strCells in chessboard represents correct numeric position on chessboard
func TestStrCellsInChessBoardRepresentsCorrectNumericPositionOnChessBoard(t *testing.T) {
	inputs := map[string]int{
		"a1": 1,
		"d5": 29,
		"e2": 34,
		"b8": 16,
		"c7": 23,
	}

	for i, v := range inputs {
		if cb.strCells[i] != v {
			t.Fatalf("Failed to Test if strCells in chessboard represents correct numeric position on chessboard. For cell %v, it represents %v numeric position. Expected output: %v ---> %v", cb.strCells[i], v, i, v)
		}
	}
}

// Test if a CreateChessboard() func returns pointer to the Chessboard type
func TestCreateChessBoardReturnsPointerToChessboardType(t *testing.T) {
	str := reflect.TypeOf(cb).String()

	if str != "*main.ChessBoard" {
		t.Error("Failed to create Chessboard!")
	}
}

// Test if a placePiece() func places piece on a chessboard
func TestPlacePieceFuncPlacesPieceOnAChessBoard(t *testing.T) {
	p = CreatePiece(20, "pawn") // Create a piece
	cb.placePiece(20, p)

	if cb.pieces[20].kind != "pawn" {
		t.Fatal("Failed to test if a placePiece() func places piece on a chessboard")
	}

	p = nil // Delete a piece
}

// Test if a placePiece() func adds/updates available move positions for a piece
func TestPlacePieceFuncUpdatesAvailableMovePositionsForPiece(t *testing.T) {
	p = CreatePiece(20, "pawn") // Create a piece
	cb.placePiece(20, p)

	if _, ok := p.availPos[21]; !ok {
		t.Fatal("Failed to test if a placePiece() func adds/updates available move positions for a piece")
	}

	p = nil // Delete a piece
}

// Benchmark test for CreateChessBoard() func
func BenchmarkCreateChessBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateChessBoard()
	}
}
