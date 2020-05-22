package chess

import (
	"testing"
)

// Test if correct move position is generated for pawn type piece
func TestCorrectMovePositionIsGeneratedForPawnTypePiece(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(20, "pawn") // Create piece
	p.AvailPos = make(map[int]bool)
	p.addPositions(cb, p.yBoundry.top, p.maxY, Positive, 1)

	if !p.AvailPos[21] {
		t.Error("Failed to test if correct move position is generated for pawn type piece!")
	}

	p = nil // Delete piece
}
