package main

import (
	"testing"
)

// Test if a chessboard has 64 cells in it
func TestChessBoardHas64Cells(t *testing.T) {
	if len(cb.cells) != 64 {
		t.Fatalf("Failed to test if a chessboard has 64 cells in it, Number of cells: %v", len(cb.cells))
	}
}
