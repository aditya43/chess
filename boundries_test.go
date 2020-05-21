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

// Test if a right boundry is set for a piece if it is allowed to move on X axis
func TestRightBoundryIsSetForPieceIfItCanMoveOnXAxis(t *testing.T) {
	p = CreatePiece(20, "rook")
	p.setBoundries()

	if p.xBoundry.right == 0 {
		t.Fatal("Failed to test if a right boundry is set for a piece if it is allowed to move on X axis ")
	}

	p = nil
}

// Test if a correct left boundry is set for a piece
func TestCorrectLeftBoundryIsSetForPiece(t *testing.T) {
	p = CreatePiece(20, "rook")
	p.setLeftBoundry()

	if p.xBoundry.left != 4 {
		t.Fatalf("Failed to test if a correct left boundry is set for a piece. Expected: %v. Got: %v", 20, p.xBoundry.left)
	}

	p = nil
}

// Test if a correct right boundry is set for a piece
func TestCorrectRightBoundryIsSetForPiece(t *testing.T) {
	p = CreatePiece(20, "rook")
	p.setRightBoundry()

	if p.xBoundry.right != 60 {
		t.Fatalf("Failed to test if a correct right boundry is set for a piece. Expected: %v. Got: %v", 60, p.xBoundry.right)
	}

	p = nil
}
