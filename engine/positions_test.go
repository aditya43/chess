package engine

import (
	"testing"
)

// Test if correct move positions are generated for pawn type piece
func TestCorrectMovePositionsAreGeneratedForPawnTypePiece(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(20, "pawn") // Create piece
	p.AvailPos = make(map[int]bool)
	p.addPositions(cb, p.yBoundry.top, p.maxY, Positive, 1)
	p.addPositions(cb, p.diagonalBoundry.topRight, p.maxCross, Positive, 9)
	p.addPositions(cb, p.diagonalBoundry.topLeft, p.maxCross, Negative, 7)

	if !p.AvailPos[21] {
		t.Error("Pawn doesn't have move position in upward direction!")
	}

	if !p.AvailPos[13] {
		t.Error("Missing top-left move position for a pawn piece!")
	}

	if !p.AvailPos[29] {
		t.Error("Missing top-right move position for a pawn piece!")
	}
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
}

// Test move positions are generated in right direction for a non pawn type piece
func TestMovePositionsAreGeneratedInRightDirectionForNonPawnTypePiece(t *testing.T) {
	cb := CreateChessBoard()
	p := CreatePiece(20, "rook") // Create piece
	p.AvailPos = make(map[int]bool)
	p.addPositions(cb, p.xBoundry.right, p.maxX, Positive, 1)

	if !p.AvailPos[28] && !p.AvailPos[36] && !p.AvailPos[44] && !p.AvailPos[52] && !p.AvailPos[60] {
		t.Errorf("Failed to test if positions are generated in right direction for a non pawn type piece! Got: %v", p.AvailPos)
	}
}

// Test correct move positions are generated for different types of pieces
func TestCorrectMovePositionsAreGeneratedForDifferentTypesOfPieces(t *testing.T) {
	// 29
	cb := CreateChessBoard()
	kinds := [6]string{"king", "queen", "bishop", "horse", "rook", "pawn"}
	positions := map[string][]int{
		"king":   {36, 20, 30, 28, 37, 21, 38, 22},
		"queen":  {13, 15, 50, 2, 28, 26, 45, 61, 22, 32, 37, 43, 57, 30, 53, 21, 5, 47, 36, 20, 27, 38, 8, 31, 25, 56, 11},
		"bishop": {15, 36, 43, 20, 11, 22, 47, 56, 8, 50, 57, 2, 38},
		"horse":  {35, 44, 46, 39, 14, 12, 23, 19},
		"rook":   {21, 27, 5, 53, 13, 30, 31, 28, 45, 61, 32, 26, 25, 37},
		"pawn":   {30, 22, 38},
	}

	for _, v := range kinds {
		p := CreatePiece(29, v)
		p.updateMovePositions(cb)

		if len(p.AvailPos) < 1 {
			t.Errorf("0 move positions are generated for piece type: %v at position 29(D5)", p.Kind)
		}

		if len(p.AvailPos) != len(positions[p.Kind]) {
			t.Errorf("Wrong number of positions generated for piece type %v. Expected %v positions. Got %v positions", p.Kind, len(positions[p.Kind]), len(p.AvailPos))
		}

		for _, pos := range positions[p.Kind] {
			if !p.AvailPos[pos] {
				t.Errorf("Expected position %v (%v) is not set for a piece type %v", pos, cb.Cells[pos], p.Kind)
			}
		}
	}
}

// Test if a horse is not allowed to make 2 left, 1 top move
func TestHorseNotAllowedToMake2Left1TopMove(t *testing.T) {
	if canMakeTwoLeftOneTop(15) {
		t.Error("Horse shoudln't make make 2 left, 1 top move")
	}

	if canMakeTwoLeftOneTop(32) {
		t.Error("Horse shoudln't make make 2 left, 1 top move")
	}
}

// Test if a horse is not allowed to make 2 left, 1 down
func TestHorseNotAllowedToMake2Left1DownMove(t *testing.T) {
	if canMakeTwoLeftOneDown(16) {
		t.Error("Horse shoudln't make make 2 left, 1 down move")
	}

	if canMakeTwoLeftOneDown(41) {
		t.Error("Horse shoudln't make make 2 left, 1 down move")
	}
}

// Test if a horse is not allowed to make 2 down, 1 right
func TestHorseNotAllowedToMake2Down1RightMove(t *testing.T) {
	if canMakeTwoDownOneRight(57) {
		t.Error("Horse shoudln't make make 2 down, 1 right move")
	}

	if canMakeTwoDownOneRight(10) {
		t.Error("Horse shoudln't make make 2 down, 1 right move")
	}
}

// Test if a horse is not allowed to make 2 top, 1 left
func TestHorseNotAllowedToMake2Top1LeftMove(t *testing.T) {
	if canMakeTwoTopOneLeft(7) {
		t.Error("Horse shoudln't make make 2 top, 1 left move")
	}

	if canMakeTwoTopOneLeft(31) {
		t.Error("Horse shoudln't make make 2 top, 1 left move")
	}
}

// Test if a horse is not allowed to make 2 down, 1 left
func TestHorseNotAllowedToMake2Down1LeftMove(t *testing.T) {
	if canMakeTwoDownOneLeft(8) {
		t.Error("Horse shoudln't make make 2 down, 1 left")
	}

	if canMakeTwoDownOneLeft(34) {
		t.Error("Horse shoudln't make make 2 down, 1 left")
	}
}

// Test if a horse is not allowed to make 2 right, 1 down
func TestHorseNotAllowedToMake2Right1DownMove(t *testing.T) {
	if canMakeTwoRightOneDown(50) {
		t.Error("Horse shoudln't make make 2 right, 1 down")
	}

	if canMakeTwoRightOneDown(17) {
		t.Error("Horse shoudln't make make 2 right, 1 down")
	}
}
