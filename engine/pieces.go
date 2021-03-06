package engine

// Chessboard piece
type Piece struct {
	Kind     string       // Kind of piece. e.g. king, queen, horse, pawn, rook, bishop
	CurPos   int          // Numeric position on board
	Symbol   string       // e.g. ♚, ♛, ♜, ♝, ♞
	AvailPos map[int]bool // Available move positions for a piece

	isHorse bool // Is current piece a Horse
	isPawn  bool // Is current piece a Pawn

	xAllow     bool // Is piece allowed to traverse in X direction
	yAllow     bool // Is piece allowed to traverse in Y direction
	crossAllow bool // Is piece allowed to traverse in diagonal direction

	maxX     int // Max number of cells a piece can traverse in X direction
	maxY     int // Max number of cells a piece can traverse in Y direction
	maxCross int // Max number of cells a piece can traverse in diagonal direction

	xBoundry        XBoundry        // Horizontal boundries
	yBoundry        YBoundry        // Vertical boundries
	diagonalBoundry DiagonalBoundry // Diagonal boundries
}

// Calculate and add available positions for a piece based off it's current position
func (p *Piece) updateMovePositions(b *ChessBoard) {
	p.AvailPos = make(map[int]bool)

	if !p.isHorse {
		if p.yAllow { // Check if a piece can move on y axis
			if p.isPawn {
				// Y axis
				// Get positions in top direction for a pawn piece
				p.addPositions(b, p.yBoundry.top, p.maxY, Positive, 1)
			}

			if !p.isPawn {
				// Y axis
				// Get positions in top direction for a non pawn piece
				p.addPositions(b, p.yBoundry.top, p.maxY, Positive, 1)

				// Y axis
				// Get positions in bottom direction for a non pawn piece
				p.addPositions(b, p.yBoundry.bottom, p.maxY, Negative, 1)

			}
		}

		// Check if a piece can move on x axis
		if p.xAllow {
			// X axis
			// Get positions in right direction for a non pawn piece
			p.addPositions(b, p.xBoundry.right, p.maxX, Positive, 8)

			// X axis
			// Get positions in left direction for a non pawn piece
			p.addPositions(b, p.xBoundry.left, p.maxX, Negative, 8)
		}

		// Check if a piece can move in diagonal direction
		if p.crossAllow {
			// Diagonal
			// Get diagonal move positions in top-right direction for a non pawn piece
			p.addPositions(b, p.diagonalBoundry.topRight, p.maxCross, Positive, 9)

			// Diagonal
			// Get diagonal move positions in top-left direction for a non pawn piece
			p.addPositions(b, p.diagonalBoundry.topLeft, p.maxCross, Negative, 7)

			// Diagonal
			// Get diagonal move positions in bottom-right direction for a non pawn piece
			p.addPositions(b, p.diagonalBoundry.bottomRight, p.maxCross, Positive, 7)

			// Diagonal
			// Get diagonal move positions in bottom-left direction for a non pawn piece
			p.addPositions(b, p.diagonalBoundry.bottomLeft, p.maxCross, Negative, 9)
		}
	}

	if p.isHorse {
		p.addHorsePositions(b) // Add move positions for a horse piece
	}
}

// Create piece
//
// pos: Numeric position of a piece on board
// n: Piece type
func CreatePiece(pos int, n string) *Piece {
	var p *Piece

	switch n {
	case "rook":
		p = &Piece{
			Kind:       n,
			CurPos:     pos,
			Symbol:     "♖",
			xAllow:     true,
			yAllow:     true,
			crossAllow: false,
			maxX:       8,
			maxY:       8,
		}
	case "horse":
		p = &Piece{
			Kind:       n,
			CurPos:     pos,
			Symbol:     "♘",
			isHorse:    true,
			xAllow:     false,
			yAllow:     false,
			crossAllow: false,
		}
	case "bishop":
		p = &Piece{
			Kind:       n,
			CurPos:     pos,
			Symbol:     "♗",
			xAllow:     false,
			yAllow:     false,
			crossAllow: true,
			maxCross:   8,
		}
	case "queen":
		p = &Piece{
			Kind:       n,
			CurPos:     pos,
			Symbol:     "♕",
			xAllow:     true,
			yAllow:     true,
			crossAllow: true,
			maxX:       8,
			maxY:       8,
			maxCross:   8,
		}
	case "king":
		p = &Piece{
			Kind:       n,
			CurPos:     pos,
			Symbol:     "♔",
			xAllow:     true,
			yAllow:     true,
			crossAllow: true,
			maxX:       1,
			maxY:       1,
			maxCross:   1,
		}
	case "pawn":
		p = &Piece{
			Kind:       n,
			CurPos:     pos,
			Symbol:     "♙",
			isPawn:     true,
			yAllow:     true,
			maxY:       1,
			crossAllow: true,
			maxCross:   1,
		}
	}

	// Set boundries for a piece if it can move on X, Y axis as well as diagonally
	p.setBoundries()

	return p
}
