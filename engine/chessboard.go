package engine

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

type ChessBoard struct {
	Cells    map[int]string // Numeric Position to Text Representation. For e.g. 10 -> B2
	Pieces   map[int]*Piece // Numeric Position --> Piece. For printing purposes only
	StrCells map[string]int // Text Representation to Numeric Position of a cell
}

// Create new chessboard
// Create cells with their numeric identifiers
func CreateChessBoard() *ChessBoard {
	b := &ChessBoard{}
	b.Cells = make(map[int]string)
	b.StrCells = make(map[string]int)

	l := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"} // Columns
	num := 1

	for _, c := range l {
		for i := 1; i < 9; i++ {
			str := fmt.Sprintf("%v%v", c, i) // Generate string position e.g. a1, h8 etc..
			b.Cells[num] = str
			b.StrCells[str] = num
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
func (b *ChessBoard) PlacePiece(pos int, p *Piece) {
	b.Pieces = make(map[int]*Piece)
	b.Pieces[pos] = p
	p.updateMovePositions(b)
}

// Render chessboard with a piece
func (b *ChessBoard) Print(p *Piece) {
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
		if p.CurPos == cell {
			s := color.New(color.FgWhite, color.BgRed, color.OpBold).Sprintf(" %v ", p.Symbol)
			// fmt.Printf("%6v %v ", b.Cells[cell], s)
			color.Red.Printf("%6v", strings.ToUpper(b.Cells[cell]))
			fmt.Printf(" %v ", s)
			// color.BgRed.Printf("%6v  %v  ", b.Cells[cell], p.Symbol)
		}

		// If a piece can move to this cell, print it in green
		if _, ok := p.AvailPos[cell]; ok {
			// color.Green.Printf("%6v (%2v)", strings.ToUpper(b.Cells[cell]), cell)
			color.Green.Printf("%6v  %2v ", strings.ToUpper(b.Cells[cell]), "")
		}

		// Piece is not at this cell and neither can move to this location
		if _, ok := p.AvailPos[cell]; !ok && p.CurPos != cell {
			// fmt.Printf("%6v (%2v)", strings.ToUpper(b.Cells[cell]), cell)
			fmt.Printf("%6v  %2v ", strings.ToUpper(b.Cells[cell]), "")
		}

		cell += 8
		cnt++
	}
	fmt.Println()
}
