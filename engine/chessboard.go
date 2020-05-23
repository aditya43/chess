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
	b.Pieces = make(map[int]*Piece)

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
	b.Pieces[pos] = p
	p.updateMovePositions(b)
}

// Render chessboard
func (b *ChessBoard) Print() {
	g := make(map[int]bool) // Green cells bucket.
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

		// Print a cell with a piece symbol
		if _, ok := b.Pieces[cell]; ok {
			s := color.New(color.FgWhite, color.BgRed, color.OpBold).Sprintf(" %v ", b.Pieces[cell].Symbol)
			color.Red.Printf("%6v", strings.ToUpper(b.Cells[cell]))
			fmt.Printf(" %v ", s)
		}

		// If any piece can move to this cell, print it in green
		if b.Pieces[cell] == nil {
			for _, v := range b.Pieces {
				if _, ok := v.AvailPos[cell]; ok {
					g[cell] = true // Add this cell to green cells bucket
					color.Green.Printf("%6v  %2v ", strings.ToUpper(b.Cells[cell]), "")
					break
				}
			}
		}

		// No piece can move here and there is no piece at this location
		if b.Pieces[cell] == nil && !g[cell] {
			fmt.Printf("%6v  %2v ", strings.ToUpper(b.Cells[cell]), "")
		}

		cell += 8
		cnt++
	}
	fmt.Println()
}
