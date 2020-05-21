package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color" // For printing colorful text in terminal
)

var cb *ChessBoard
var p *Piece
var cr *bufio.Reader

// Create chessboard
func init() {
	cb = CreateChessBoard()
}

// Execute program
func main() {
	for {
		cr = bufio.NewReader(os.Stdin)
		printInfo()
		s, _ := cr.ReadString('\n')
		w := strings.Fields(strings.ToLower(s))
		checkExit(s)
		takeAction(w)
	}
}

// Take action based off user input
func takeAction(w []string) {
	err := validateInput(w)

	if err != nil {
		color.Error.Println(err)
		restartProgram()
		return
	}

	pos := cb.strCells[w[1]] // Get numeric position for a cell

	p = CreatePiece(pos, w[0]) // Create a piece
	cb.placePiece(pos, p)      // Place a piece on chessboard
	printOutput()
	restartProgram()
}

// Validate user inputs
func validateInput(w []string) error {
	var pieceErr, posErr bool

	if len(w) < 2 {
		return errors.New("Invalid input. Please enter piece type and position.")
	}

	switch strings.ToLower(w[0]) {
	case "king", "queen", "bishop", "horse", "rook", "pawn":
	default:
		pieceErr = true // User didn't input any of the above piece type
	}

	switch strings.ToLower(w[1]) {
	case "a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8",
		"b1", "b2", "b3", "b4", "b5", "b6", "b7", "b8",
		"c1", "c2", "c3", "c4", "c5", "c6", "c7", "c8",
		"d1", "d2", "d3", "d4", "d5", "d6", "d7", "d8",
		"e1", "e2", "e3", "e4", "e5", "e6", "e7", "e8",
		"f1", "f2", "f3", "f4", "f5", "f6", "f7", "f8",
		"g1", "g2", "g3", "g4", "g5", "g6", "g7", "g8",
		"h1", "h2", "h3", "h4", "h5", "h6", "h7", "h8":
	default:
		posErr = true // User didn't input any of the above cell positions
	}

	if pieceErr && posErr {
		return errors.New("Invalid piece type and position value.")
	}

	if pieceErr {
		return errors.New("Invalid piece type.")
	}

	if posErr {
		return errors.New("Invalid position value.")
	}

	return nil
}

// Print information about how to use this program
func printInfo() {
	print("\033[H\033[2J") // Ansi escape sequence to clear the screen
	fmt.Println("Enter a piece type and it's position")
	fmt.Println("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
	color.New(color.FgGreen, color.BgBlack, color.OpBold).Println("👍 Program input is case-insensitive.")
	color.New(color.FgGreen, color.BgBlack, color.OpBold).Print("👍 Type ")
	color.BgRed.Print("exit")
	color.New(color.FgGreen, color.BgBlack, color.OpBold).Println(" to quit program.")
	fmt.Print("\n")

	fmt.Print("Piece types: ")
	color.Yellow.Println("King, Queen, Bishop, Horse, Rook, Pawn")
	fmt.Print("Example input: ")
	color.Style{color.FgGreen, color.OpBold}.Println("King D5")

	fmt.Println("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
	fmt.Print("Input: ")
}

// Exit program if user typed 'exit'
func checkExit(s string) {
	if strings.HasPrefix(strings.ToLower(s), "exit") {
		fmt.Println("Good bye!🍺")
		os.Exit(0)
	}
}

// Ask user to hit Enter key to restart program
func restartProgram() {
	p = nil // Delete piece
	color.Yellow.Println("Hit Enter key to continue")
	s, _ := cr.ReadString('\n')
	checkExit(s) // User can quit program by typing 'exit'
}

// Print available move positions for a piece on chaseboard
func printOutput() {
	m := ""
	color.Yellow.Print("Available Moves: ")

	for pos := range p.availPos {
		m += strings.ToUpper(cb.cells[pos]) + ", "
	}
	color.New(color.FgGreen, color.BgBlack, color.OpBold).Printf("%v\n\n\n", strings.TrimSuffix(m, ", "))
	cb.print(p)
	fmt.Print("\n")
}
