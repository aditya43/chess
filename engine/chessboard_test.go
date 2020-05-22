package engine

import (
	"reflect"
	"testing"
)

// Test if a chessboard has 64 cells in it
func TestChessBoardHas64Cells(t *testing.T) {
	cb := CreateChessBoard()
	if len(cb.Cells) != 64 {
		t.Fatalf("Failed to test if a chessboard has 64 cells in it, Number of cells: %v", len(cb.Cells))
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

	cb := CreateChessBoard()

	for i, v := range inputs {
		if cb.StrCells[i] != v {
			t.Fatalf("Failed to Test if strCells in chessboard represents correct numeric position on chessboard. For cell %v, it represents %v numeric position. Expected output: %v ---> %v", cb.StrCells[i], v, i, v)
		}
	}
}

// Test if a CreateChessboard() func returns pointer to the Chessboard type
func TestCreateChessBoardReturnsPointerToChessboardType(t *testing.T) {
	cb := CreateChessBoard()
	str := reflect.TypeOf(cb).String()

	if str != "*chess.ChessBoard" {
		t.Error("Failed to create Chessboard!")
	}
}

// Test if a placePiece() func places piece on a chessboard
func TestPlacePieceFuncPlacesPieceOnAChessBoard(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(20, "pawn") // Create a piece
	cb.PlacePiece(20, p)

	if cb.Pieces[20].Kind != "pawn" {
		t.Fatal("Failed to test if a placePiece() func places piece on a chessboard")
	}
}

// Test if a placePiece() func adds/updates available move positions for a piece
func TestPlacePieceFuncUpdatesAvailableMovePositionsForPiece(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(20, "pawn") // Create a piece
	cb.PlacePiece(20, p)

	if _, ok := p.AvailPos[21]; !ok {
		t.Fatal("Failed to test if a placePiece() func adds/updates available move positions for a piece")
	}
}
