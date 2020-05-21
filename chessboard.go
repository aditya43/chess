package main

import "fmt"

type ChessBoard struct {
	cells    map[int]string // Numeric Position to Text Representation. For e.g. 10 -> B2
	pieces   map[int]*Piece // Numeric Position --> Piece. For printing purposes only
	strCells map[string]int // Text Representation to Numeric Position of a cell
}

// Create new chessboard
// Create cells with their numeric identifiers
func CreateChessBoard() *ChessBoard {
	b := &ChessBoard{}
	b.cells = make(map[int]string)
	b.strCells = make(map[string]int)

	l := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"} // Columns
	num := 1

	for _, c := range l {
		for i := 1; i < 9; i++ {
			str := fmt.Sprintf("%v%v", c, i) // Generate string position e.g. a1, h8 etc..
			b.cells[num] = str
			b.strCells[str] = num
			num++
		}
	}

	return b
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
