package main

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

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
	b.pieces = make(map[int]*Piece)
	b.pieces[pos] = p
	p.updateMovePositions(b)
}

// Render chessboard with a piece
func (b *ChessBoard) print(p *Piece) {
	row := 8
	cnt := 0
	cell := 8
	for i := 0; i < 64; i++ {
		if cnt == 8 {
			cnt = 0
			row -= 1
			cell = row
			fmt.Println()
		}

		// Print piece symbol if we are at current position of a piece
		if p.curPos == cell {
			s := color.New(color.FgWhite, color.BgRed, color.OpBold).Sprintf(" %v ", p.symbol)
			// fmt.Printf("%6v %v ", b.cells[cell], s)
			color.Red.Printf("%6v", strings.ToUpper(b.cells[cell]))
			fmt.Printf(" %v ", s)
			// color.BgRed.Printf("%6v  %v  ", b.cells[cell], p.symbol)
		}

		// If a piece can move to this cell, print it in green
		if _, ok := p.availPos[cell]; ok {
			// color.Green.Printf("%6v (%2v)", strings.ToUpper(b.cells[cell]), cell)
			color.Green.Printf("%6v  %2v ", strings.ToUpper(b.cells[cell]), "")
		}

		// Piece is not at this cell and neither can move to this location
		if _, ok := p.availPos[cell]; !ok && p.curPos != cell {
			// fmt.Printf("%6v (%2v)", strings.ToUpper(b.cells[cell]), cell)
			fmt.Printf("%6v  %2v ", strings.ToUpper(b.cells[cell]), "")
		}

		cell += 8
		cnt++
	}
	fmt.Println()
}
