package chess

import (
	"reflect"
	"testing"
)

// Test if a CreatePiece() func returns pointer to the Piece type
func TestCreatePieceReturnsPointerToPieceType(t *testing.T) {
	p := CreatePiece(20, "pawn") // Create piece
	str := reflect.TypeOf(p).String()

	if str != "*chess.Piece" {
		t.Error("Failed to test if a CreatePiece() func returns pointer to the Piece type!")
	}
}

// Test if a CreatePiece() func creates a piece
func TestCreatePieceCreatesPiece(t *testing.T) {
	p := CreatePiece(20, "pawn") // Create piece

	if p.Kind != "pawn" {
		t.Error("Failed to test if a CreatePiece() func creates a piece")
	}
}

// Test if the updateMovePositions() func updates available move positions for a piece
func TestUpdateMovePositionsUpdatesAvailableMovePositionsForPiece(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(20, "pawn") // Create piece
	p.updateMovePositions(cb)

	if len(p.AvailPos) < 1 {
		t.Error("Failed to test if a updateMovePositions() func updates available move positions for a piece")
	}
}

// Test if the updateMovePositions() func updates correct number of available move positions for a piece
func TestUpdateMovePositionsUpdateCorrectNumberOfMovePositionsForPiece(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(29, "horse") // Create piece
	p.updateMovePositions(cb)

	if len(p.AvailPos) != 8 {
		t.Error("Failed to test if a updateMovePositions() func updates correct number of available move positions for a piece")
	}
}
