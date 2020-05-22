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

// Test if the updateMovePositions() func generates correct move positions for a piece
func TestUpdateMovePositionsGenerateCorrectMovePositionsForPiece(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(29, "horse") // Create piece
	p.updateMovePositions(cb)

	correctMoves := [8]int{39, 46, 44, 35, 12, 19, 14, 23}

	for _, v := range correctMoves {
		if _, ok := p.AvailPos[v]; !ok {
			t.Errorf("Failed to test if the updateMovePositions() func generates correct move positions for a piece. Found a missing move position for a piece: %v ---> %v", v, cb.Cells[v])
		}
	}
}

// Test if move positions are generated for different types of pieces
func TestMovePositionsAreGeneratedForDifferentTypesOfPieces(t *testing.T) {
	kinds := [6]string{"king", "queen", "bishop", "horse", "rook", "pawn"}

	for _, v := range kinds {
		p := CreatePiece(28, v)

		if len(p.AvailPos) < 0 {
			t.Fatalf("Move positions are not generated for piece type %v at position %v", p.Kind, p.CurPos)
		}

		p = nil
	}
}
