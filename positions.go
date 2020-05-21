package main

type Direction int // Move direction of a piece

const (
	Positive Direction = iota
	Negative
)

// Calculate and add available move positions for any piece except horse
//
// b 	ChessBoard
// l 	Limit
// max  Maximum distance a piece can travel
// d 	Positive or Negative direction
// f 	Increment/Decrement factor
func (p *Piece) addPositions(b *ChessBoard, l int, max int, d Direction, f int) {
	if l < 1 {
		return // Piece doesn't have any space to move in d Direction
	}

	pos := p.curPos
	for i := 0; i < max; i++ {
		if l > 0 && pos == l { // Position has reached boundry
			if pos != p.curPos {
				p.availPos[pos] = true // Add move position
			}
			break // Reached boundry, break!
		}

		// For calculating positive move positions
		if d == Positive {
			pos += f // Increment by factor
		}

		// For calculating negative move positions
		if d == Negative {
			pos -= f // Decrement by factor
		}

		p.availPos[pos] = true // Add move position
	}
}

// Add move position for a horse
func (p *Piece) addHorsePositions(b *ChessBoard) {
	// 2 left, 1 top move

	// 2 left, 1 down move

	// 2 top, 1 left move

	// 2 down, 1 left move

	// 2 down, 1 right move

	// 2 right, 1 down move

	// 2 right, 1 top move

	// 2 top, 1 right move

}

// Can a piece move 2 left, 1 top
// p 	current position of a piece
func canMakeTwoLeftOneTop(p int) bool {
	if p <= 16 {
		return false
	}

	switch p {
	case 24, 32, 40, 48, 56, 64:
		return false
	}
	return true
}
