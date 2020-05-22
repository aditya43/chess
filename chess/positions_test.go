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

// Test move positions are generated in top direction for a non pawn type piece
func TestMovePositionsAreGeneratedInTopDirectionForNonPawnTypePiece(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(20, "rook") // Create piece
	p.AvailPos = make(map[int]bool)
	p.addPositions(cb, p.yBoundry.top, p.maxY, Positive, 1)

	if !p.AvailPos[21] && !p.AvailPos[22] && !p.AvailPos[23] && !p.AvailPos[24] {
		t.Errorf("Failed to test if positions are generated in top direction for a non pawn type piece! Got: %v", p.AvailPos)
	}

	p = nil // Delete piece
}

// Test move positions are generated in bottom direction for a non pawn type piece
func TestMovePositionsAreGeneratedInBottomDirectionForNonPawnTypePiece(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(20, "rook") // Create piece
	p.AvailPos = make(map[int]bool)
	p.addPositions(cb, p.yBoundry.bottom, p.maxY, Negative, 1)

	if !p.AvailPos[19] && !p.AvailPos[18] && !p.AvailPos[17] {
		t.Errorf("Failed to test if positions are generated in bottom direction for a non pawn type piece! Got: %v", p.AvailPos)
	}

	p = nil // Delete piece
}

// Test move positions are generated in left direction for a non pawn type piece
func TestMovePositionsAreGeneratedInLeftDirectionForNonPawnTypePiece(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(20, "rook") // Create piece
	p.AvailPos = make(map[int]bool)
	p.addPositions(cb, p.xBoundry.left, p.maxX, Negative, 1)

	if !p.AvailPos[12] && !p.AvailPos[4] {
		t.Errorf("Failed to test if positions are generated in left direction for a non pawn type piece! Got: %v", p.AvailPos)
	}

	p = nil // Delete piece
}
