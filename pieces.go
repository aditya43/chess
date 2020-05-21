package main

type Piece struct {
	name       string       // Name. e.g. king, queen, horse, pawn, rook, bishop
	curPos     int          // Numeric position on board
	symbol     string       // e.g. ♚, ♛, ♜, ♝, ♞
	isHorse    bool         // Is current piece a Horse
	isPawn     bool         // Is current piece a Pawn
	xAllow     bool         // Is piece allowed to traverse in X direction
	yAllow     bool         // Is piece allowed to traverse in Y direction
	crossAllow bool         // Is piece allowed to traverse in diagonal direction
	maxX       int          // Max number of cells a piece can traverse in X direction
	maxY       int          // Max number of cells a piece can traverse in Y direction
	maxCross   int          // Max number of cells a piece can traverse in diagonal direction
	availPos   map[int]bool // Available move positions for a piece
}
