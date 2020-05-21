package main

type ChessBoard struct {
	cells    map[int]string // Numeric Position to Text Representation. For e.g. 10 -> B2
	pieces   map[int]*Piece // Numeric Position --> Piece. For printing purposes only
	strCells map[string]int // Text Representation to Numeric Position of a cell
}

// Create new chessboard
// Create cells with their numeric identifiers
func CreateChessBoard() *ChessBoard {
	//
}

// Place piece on a chessboard and update
// available move positions for a piece.
//
// pos 	numeric position on board
// p 	Piece
func (b *ChessBoard) placePiece(pos int, p *Piece) {
	//
}

// Render chessboard with a piece
func (b *ChessBoard) print(p *Piece) {
	//
}
