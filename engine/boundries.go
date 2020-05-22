package engine

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
	pos := p.CurPos

	for i := 0; i < 8; i++ {
		if pos <= 1 {
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
	pos := p.CurPos

	for i := 0; i < 8; i++ {
		if pos >= 64 {
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
	pos := p.CurPos

	for i := 0; i < 8; i++ {
		if pos >= 64 {
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
	pos := p.CurPos

	for i := 0; i < 8; i++ {
		if pos <= 1 {
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
	pos := p.CurPos

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
				if p.isPawn {
					// Pawn is not on left boundry or top boundry
					// Calculate and set its top-left position
					p.diagonalBoundry.topLeft = pos - 7
					return
				}
				pos -= 7
			}
		}
	}
}

// Calculate top-right boundry based off current position of a piece
func (p *Piece) setTopRightBoundry() {
	pos := p.CurPos

	for i := 0; i < 8; i++ {
		if pos >= 64 {
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
				if p.isPawn {
					// Pawn is not on right boundry or top boundry
					// Calculate and set its top-right position
					p.diagonalBoundry.topRight = pos + 9
					return
				}
				pos += 9
			}
		}
	}
}

// Calculate bottom-right boundry based off current position of a piece
func (p *Piece) setBottomRightBoundry() {
	if p.isPawn {
		return
	}

	pos := p.CurPos

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
	if p.isPawn {
		return
	}

	pos := p.CurPos

	for i := 0; i < 8; i++ {
		if pos <= 1 {
			return
		}

		if pos >= 1 {
			switch pos {
			case 1, 9, 17, 25, 33, 41, 49, 57: // Bottom boundry numeric positions on board
				p.diagonalBoundry.bottomLeft = pos
				return
			case 2, 3, 4, 5, 6, 7, 8: // Left boundry numeric positions on board
				p.diagonalBoundry.bottomLeft = pos
				return
			default:
				pos -= 9
			}
		}
	}
}

// Set boundries for a piece
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

	// fmt.Printf("%v %8v: %v %v %v %v %v %v %v %v \n", p.Symbol, p.Kind, p.xBoundry.left, p.diagonalBoundry.topLeft, p.yBoundry.top, p.diagonalBoundry.topRight, p.xBoundry.right, p.diagonalBoundry.bottomRight, p.yBoundry.bottom, p.diagonalBoundry.bottomLeft)
}
