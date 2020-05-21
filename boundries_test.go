package main

import (
	"testing"
)

// Test if a left boundry is set for a piece if it is allowed to move on X axis
func TestLeftBoundryIsSetForPieceIfItCanMoveOnXAxis(t *testing.T) {
	p = CreatePiece(20, "rook")
	p.setBoundries()

	if p.xBoundry.left == 0 {
		t.Fatal("Failed to test if a left boundry is set for a piece if it is allowed to move on X axis ")
	}

	p = nil
}
