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
