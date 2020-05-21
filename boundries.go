package main

// Boundries in horizontal direction for a piece
type XBoundry struct {
	left  int
	right int
}

// Boundries in vertical direction for a piece
type YBoundry struct {
	top    int
	bottom int
}

// Diagonal boundries for a piece
type DiagonalBoundry struct {
	topLeft     int
	topRight    int
	bottomLeft  int
	bottomRight int
}

// Calculate left boundry based off current position of a piece
func (p *Piece) setLeftBoundry() {
	pos := p.curPos

	for i := 0; i < 8; i++ {
		if pos < 1 {
			return
		}

		if pos >= 1 {
			switch pos {
			case 1, 2, 3, 4, 5, 6, 7, 8: // Left boundry numeric positions on board
				p.xBoundry.left = pos
				return
			default:
				pos -= 8
			}
		}
	}
}

// Calculate right boundry based off current position of a piece
func (p *Piece) setRightBoundry() {
	pos := p.curPos

	for i := 0; i < 8; i++ {
		if pos > 64 {
			return
		}

		if pos <= 64 {
			switch pos {
			case 57, 58, 59, 60, 61, 62, 63, 64: // Right boundry numeric positions on board
				p.xBoundry.right = pos
				return
			default:
				pos += 8
			}
		}
	}
}

// Calculate top boundry based off current position of a piece
func (p *Piece) setTopBoundry() {
	pos := p.curPos

	for i := 0; i < 8; i++ {
		if pos > 64 {
			return
		}

		if pos <= 64 {
			switch pos {
			case 8, 16, 24, 32, 40, 48, 56, 64: // Top boundry numeric positions on board
				p.yBoundry.top = pos
				return
			default:
				pos++
			}
		}
	}
}

// Calculate bottom boundry based off current position of a piece
func (p *Piece) setBottomBoundry() {
	pos := p.curPos

	for i := 0; i < 8; i++ {
		if pos < 1 {
			return
		}

		if pos >= 1 {
			switch pos {
			case 1, 9, 17, 25, 33, 41, 49, 57: // Bottom boundry numeric positions on board
				p.yBoundry.bottom = pos
				return
			default:
				pos--
			}
		}
	}
}

// Calculate top-left boundry based off current position of a piece
func (p *Piece) setTopLeftBoundry() {
	pos := p.curPos

	for i := 0; i < 8; i++ {
		// fmt.Println(pos)
		if pos < 2 {
			return
		}

		if pos >= 1 {
			switch pos {
			case 8, 16, 24, 32, 40, 48, 56, 64: // Top boundry numeric positions on board
				p.diagonalBoundry.topLeft = pos
				return
			case 1, 2, 3, 4, 5, 6, 7: // Left boundry numeric positions on board
				p.diagonalBoundry.topLeft = pos
				return
			default:
				pos -= 7
			}
		}
	}
}

// Calculate top-right boundry based off current position of a piece
func (p *Piece) setTopRightBoundry() {
	pos := p.curPos

	for i := 0; i < 8; i++ {
		if pos > 64 {
			return
		}

		if pos <= 64 {
			switch pos {
			case 8, 16, 24, 32, 40, 48, 56: // Top boundry numeric positions on board
				p.diagonalBoundry.topRight = pos
				return
			case 57, 58, 59, 60, 61, 62, 63, 64: // Right boundry numeric positions on board
				p.diagonalBoundry.topRight = pos
				return
			default:
				pos += 9
			}
		}
	}
}

// Calculate bottom-right boundry based off current position of a piece
func (p *Piece) setBottomRightBoundry() {
	pos := p.curPos

	for i := 0; i < 8; i++ {
		if pos < 9 {
			return
		}

		if pos >= 9 {
			switch pos {
			case 1, 9, 17, 25, 33, 41, 49, 57: // Bottom boundry numeric positions on board
				p.diagonalBoundry.bottomRight = pos
				return
			case 58, 59, 60, 61, 62, 63, 64: // Right boundry numeric positions on board
				p.diagonalBoundry.bottomRight = pos
				return
			default:
				pos += 7
			}
		}
	}
}

// Calculate bottom-left boundry based off current position of a piece
func (p *Piece) setBottomLeftBoundry() {
	//
}

// Set boundries for a piece if it can move on X,
// Y axis as well as in diagonal direction
func (p *Piece) setBoundries() {
	if p.xAllow {
		p.setLeftBoundry()
		p.setRightBoundry()
	}

	if p.yAllow {
		p.setTopBoundry()
		p.setBottomBoundry()
	}

	if p.crossAllow {
		p.setTopLeftBoundry()
		p.setTopRightBoundry()
		p.setBottomRightBoundry()
		p.setBottomLeftBoundry()
	}

	// fmt.Printf("%v %8v: %v %v %v %v %v %v %v %v \n", p.symbol, p.name, p.xBoundry.left, p.diagonalBoundry.topLeft, p.yBoundry.top, p.diagonalBoundry.topRight, p.xBoundry.right, p.diagonalBoundry.bottomRight, p.yBoundry.bottom, p.diagonalBoundry.bottomLeft)
}
