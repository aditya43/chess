package main

type Piece struct {
	name            string          // Name. e.g. king, queen, horse, pawn, rook, bishop
	curPos          int             // Numeric position on board
	symbol          string          // e.g. ♚, ♛, ♜, ♝, ♞
	isHorse         bool            // Is current piece a Horse
	isPawn          bool            // Is current piece a Pawn
	xAllow          bool            // Is piece allowed to traverse in X direction
	yAllow          bool            // Is piece allowed to traverse in Y direction
	crossAllow      bool            // Is piece allowed to traverse in diagonal direction
	maxX            int             // Max number of cells a piece can traverse in X direction
	maxY            int             // Max number of cells a piece can traverse in Y direction
	maxCross        int             // Max number of cells a piece can traverse in diagonal direction
	availPos        map[int]bool    // Available move positions for a piece
	xBoundry        XBoundry        // Horizontal boundries
	yBoundry        YBoundry        // Vertical boundries
	diagonalBoundry DiagonalBoundry // Diagonal boundries
}

// Calculate and add available positions for a piece based off it's current position
func (p *Piece) updateMovePositions(b *ChessBoard) {
	//
}

// Create piece
// pos: Numeric position of a piece on board
// n: Piece type
func CreatePiece(pos int, n string) *Piece {
	var p *Piece

	switch n {
	case "rook":
		p = &Piece{
			name:       n,
			curPos:     pos,
			symbol:     "♖",
			xAllow:     true,
			yAllow:     true,
			crossAllow: false,
			maxX:       8,
			maxY:       8,
		}
	case "horse":
		p = &Piece{
			name:       n,
			curPos:     pos,
			symbol:     "♘",
			isHorse:    true,
			xAllow:     false,
			yAllow:     false,
			crossAllow: false,
		}
	case "bishop":
		p = &Piece{
			name:       n,
			curPos:     pos,
			symbol:     "♗",
			xAllow:     false,
			yAllow:     false,
			crossAllow: true,
			maxCross:   8,
		}
	case "queen":
		p = &Piece{
			name:       n,
			curPos:     pos,
			symbol:     "♕",
			xAllow:     true,
			yAllow:     true,
			crossAllow: true,
			maxX:       8,
			maxY:       8,
			maxCross:   8,
		}
	case "king":
		p = &Piece{
			name:       n,
			curPos:     pos,
			symbol:     "♔",
			xAllow:     true,
			yAllow:     true,
			crossAllow: true,
			maxX:       1,
			maxY:       1,
			maxCross:   1,
		}
	case "pawn":
		p = &Piece{
			name:   n,
			curPos: pos,
			symbol: "♙",
			isPawn: true,
			yAllow: true,
			maxY:   1,
		}
	}

	// Set boundries for a piece if it can move on X, Y axis as well as diagonally
	p.setBoundries()

	return p
}
