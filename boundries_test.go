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

// Test if a left boundry is not set for a piece if it is not allowed to move on X axis
func TestLeftBoundryIsNotSetForPieceIfItCannotMoveOnXAxis(t *testing.T) {
	p = CreatePiece(20, "pawn")
	p.setBoundries()

	if p.xBoundry.left != 0 {
		t.Fatal("Failed to test if a left boundry is not set for a piece if it is not allowed to move on X axis")
	}
}

// Test if a right boundry is not set for a piece if it is not allowed to move on X axis
func TestRightBoundryIsNotSetForPieceIfItCannotMoveOnXAxis(t *testing.T) {
	p = CreatePiece(20, "pawn")
	p.setBoundries()

	if p.xBoundry.right != 0 {
		t.Fatal("Failed to test if a right boundry is not set for a piece if it is not allowed to move on X axis")
	}

	p = nil
}

// Test if a top boundry is set for a piece if it is allowed to move on Y axis
func TestTopBoundryIsSetForPieceIfItCanMoveOnYAxis(t *testing.T) {
	p = CreatePiece(20, "rook")
	p.setBoundries()

	if p.yBoundry.top == 0 {
		t.Fatal("Failed to test if a top boundry is set for a piece if it is allowed to move on Y axis ")
	}

	p = nil
}

// Test if a bottom boundry is set for a piece if it is allowed to move on Y axis
func TestBottomBoundryIsSetForPieceIfItCanMoveOnYAxis(t *testing.T) {
	p = CreatePiece(20, "rook")
	p.setBoundries()

	if p.yBoundry.bottom == 0 {
		t.Fatal("Failed to test if a bottom boundry is set for a piece if it is allowed to move on Y axis ")
	}

	p = nil
}

// Test if a correct top boundry is set for a piece
func TestCorrectTopBoundryIsSetForPiece(t *testing.T) {
	p = CreatePiece(20, "rook")
	p.setTopBoundry()

	if p.yBoundry.top != 24 {
		t.Fatalf("Failed to test if a correct top boundry is set for a piece. Expected: %v. Got: %v", 20, p.yBoundry.top)
	}

	p = nil
}

// Test if a correct bottom boundry is set for a piece
func TestCorrectBottomBoundryIsSetForPiece(t *testing.T) {
	p = CreatePiece(20, "rook")
	p.setBottomBoundry()

	if p.yBoundry.bottom != 17 {
		t.Fatalf("Failed to test if a correct bottom boundry is set for a piece. Expected: %v. Got: %v", 60, p.yBoundry.bottom)
	}

	p = nil
}

// Test if a top boundry is not set for a piece if it is not allowed to move on Y axis
func TestTopBoundryIsNotSetForPieceIfItCannotMoveOnYAxis(t *testing.T) {
	p = CreatePiece(20, "bishop")
	p.setBoundries()

	if p.yBoundry.top != 0 {
		t.Fatal("Failed to test if a top boundry is not set for a piece if it is not allowed to move on Y axis")
	}

	p = nil
}

// Test if a bottom boundry is not set for a piece if it is not allowed to move on Y axis
func TestBottomBoundryIsNotSetForPieceIfItCannotMoveOnYAxis(t *testing.T) {
	p = CreatePiece(20, "bishop")
	p.setBoundries()

	if p.yBoundry.bottom != 0 {
		t.Fatal("Failed to test if a bottom boundry is not set for a piece if it is not allowed to move on Y axis")
	}

	p = nil
}

// Test if a top-left boundry is set for a piece if it is allowed to move in diagonal direction
func TestTopLeftBoundryIsSetForPieceIfItCanMoveInDiagonalDirection(t *testing.T) {
	p = CreatePiece(20, "bishop")
	p.setBoundries()

	if p.diagonalBoundry.topLeft == 0 {
		t.Fatal("Failed to test if a top-left boundry is set for a piece if it is allowed to move in diagonal direction ")
	}

	p = nil
}

// Test if a top-right boundry is set for a piece if it is allowed to move in diagonal direction
func TestTopRightBoundryIsSetForPieceIfItCanMoveInDiagonalDirection(t *testing.T) {
	p = CreatePiece(20, "bishop")
	p.setBoundries()

	if p.diagonalBoundry.topRight == 0 {
		t.Fatal("Failed to test if a top-right boundry is set for a piece if it is allowed to move in diagonal direction ")
	}

	p = nil
}

// Test if a correct top-left boundry is set for a piece
func TestCorrectTopLeftBoundryIsSetForPiece(t *testing.T) {
	p = CreatePiece(20, "bishop")
	p.setTopLeftBoundry()

	if p.diagonalBoundry.topLeft != 6 {
		t.Fatalf("Failed to test if a correct top-left boundry is set for a piece. Expected: %v. Got: %v", 20, p.diagonalBoundry.topLeft)
	}

	p = nil
}

// Test if a correct top-right boundry is set for a piece
func TestCorrectTopRightBoundryIsSetForPiece(t *testing.T) {
	p = CreatePiece(22, "bishop")
	p.setTopRightBoundry()

	if p.diagonalBoundry.topRight != 40 {
		t.Fatalf("Failed to test if a correct top-right boundry is set for a piece. Expected: %v. Got: %v", 40, p.diagonalBoundry.topRight)
	}

	p = nil
}
