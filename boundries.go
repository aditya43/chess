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
	//
}

// Calculate right boundry based off current position of a piece
func (p *Piece) setRightBoundry() {
	//
}

// Calculate top boundry based off current position of a piece
func (p *Piece) setTopBoundry() {
	//
}

// Calculate bottom boundry based off current position of a piece
func (p *Piece) setBottomBoundry() {
	//
}

// Calculate top-left boundry based off current position of a piece
func (p *Piece) setTopLeftBoundry() {
	//
}

// Calculate top-right boundry based off current position of a piece
func (p *Piece) setTopRightBoundry() {
	//
}

// Calculate bottom-right boundry based off current position of a piece
func (p *Piece) setBottomRightBoundry() {
	//
}

// Calculate bottom-left boundry based off current position of a piece
func (p *Piece) setBottomLeftBoundry() {
	//
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

	// fmt.Printf("%v %8v: %v %v %v %v %v %v %v %v \n", p.symbol, p.name, p.xBoundry.left, p.diagonalBoundry.topLeft, p.yBoundry.top, p.diagonalBoundry.topRight, p.xBoundry.right, p.diagonalBoundry.bottomRight, p.yBoundry.bottom, p.diagonalBoundry.bottomLeft)
}
