package main

type Direction int // Move direction of a piece

const (
	Positive Direction = iota
	Negative
)

// Calculate and add available move positions for a piece
//
// b 	ChessBoard
// l 	Limit
// max  Maximum distance a piece can travel
// d 	Positive or Negative direction
// f 	Increment/Decrement factor
func (p *Piece) addPositions(b *ChessBoard, l int, max int, d Direction, f int) {
	//
}
