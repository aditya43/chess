package engine

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

	pos := p.CurPos
	for i := 0; i < max; i++ {
		if l > 0 && pos == l { // Position has reached boundry
			if pos != p.CurPos {
				p.AvailPos[pos] = true // Add move position
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

		p.AvailPos[pos] = true // Add move position
	}
}

// Add move position for a horse
func (p *Piece) addHorsePositions(b *ChessBoard) {
	// 2 left, 1 top move
	if canMakeTwoLeftOneTop(p.CurPos) {
		p.AvailPos[p.CurPos-15] = true
	}

	// 2 left, 1 down move
	if canMakeTwoLeftOneDown(p.CurPos) {
		p.AvailPos[p.CurPos-17] = true
	}

	// 2 top, 1 left move
	if canMakeTwoTopOneLeft(p.CurPos) {
		p.AvailPos[p.CurPos-6] = true
	}

	// 2 down, 1 left move
	if canMakeTwoDownOneLeft(p.CurPos) {
		p.AvailPos[p.CurPos-10] = true
	}

	// 2 down, 1 right move
	if canMakeTwoDownOneRight(p.CurPos) {
		p.AvailPos[p.CurPos+6] = true
	}

	// 2 right, 1 down move
	if canMakeTwoRightOneDown(p.CurPos) {
		p.AvailPos[p.CurPos+15] = true
	}

	// 2 right, 1 top move
	if canMakeTwoRightOneTop(p.CurPos) {
		p.AvailPos[p.CurPos+17] = true
	}

	// 2 top, 1 right move
	if canMakeTwoTopOneRight(p.CurPos) {
		p.AvailPos[p.CurPos+10] = true
	}
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

// Can a piece move 2 left, 1 down
// p 	current position of a piece
func canMakeTwoLeftOneDown(p int) bool {
	if p <= 17 {
		return false
	}

	switch p {
	case 25, 33, 41, 49, 57:
		return false
	}
	return true
}

// Can a piece move 2 down, 1 right
// p 	current position of a piece
func canMakeTwoDownOneRight(p int) bool {
	if p >= 57 {
		return false
	}

	switch p {
	case 1, 2, 10, 18, 26, 34, 42, 50, 58, 9, 17, 25, 33, 41, 49:
		return false
	}
	return true
}

// Can a piece move 2 top, 1 left
// p 	current position of a piece
func canMakeTwoTopOneLeft(p int) bool {
	if p <= 8 {
		return false
	}

	switch p {
	case 16, 24, 32, 40, 48, 56, 64, 15, 23, 31, 39, 47, 55, 63:
		return false
	}
	return true
}

// Can a piece move 2 down, 1 left
// p 	current position of a piece
func canMakeTwoDownOneLeft(p int) bool {
	if p <= 10 {
		return false
	}

	switch p {
	case 10, 18, 26, 34, 42, 50, 58, 9, 17, 25, 33, 41, 49, 57:
		return false
	}
	return true
}

// Can a piece move 2 right, 1 down
// p 	current position of a piece
func canMakeTwoRightOneDown(p int) bool {
	if p >= 49 {
		return false
	}

	switch p {
	case 1, 9, 17, 25, 33, 41:
		return false
	}
	return true
}

// Can a piece move 2 right, 1 top
// p 	current position of a piece
func canMakeTwoRightOneTop(p int) bool {
	if p >= 49 {
		return false
	}

	switch p {
	case 8, 16, 24, 32, 40, 48, 56, 64:
		return false
	}
	return true
}

// Can a piece move 2 top, 1 right
// p 	current position of a piece
func canMakeTwoTopOneRight(p int) bool {
	if p >= 57 {
		return false
	}

	switch p {
	case 7, 15, 23, 31, 39, 47, 55, 63, 8, 16, 24, 32, 40, 48, 56, 64:
		return false
	}
	return true
}
