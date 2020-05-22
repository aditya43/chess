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
