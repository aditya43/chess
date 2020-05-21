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
