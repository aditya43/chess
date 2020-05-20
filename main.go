package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color" // For printing colorful text in terminal
)

var cr *bufio.Reader

func init() {
	// Create chessboard
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
		// Print error
		restartProgram()
		return
	}
	printOutput()
	restartProgram()
}

func validateInput(w []string) error {
	// Validate user inputs
	return nil
}

// Print information about how to use this program
func printInfo() {
	print("\033[H\033[2J") // Ansi escape sequence to clear the screen
	fmt.Println("Enter a piece type and it's position")
	fmt.Println("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
	color.New(color.FgGreen, color.BgBlack, color.OpBold).Println("👍 Program is not case-sensitive.")
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

func restartProgram() {
	// Ask user to hit Enter key to restart program
	// User can quit program by typing 'exit'
}

func printOutput() {
	// Print available move positions and a chaseboard
}
