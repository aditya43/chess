package main

import (
	"bufio"
	"os"
	"strings"
	// For printing colorful text in terminal
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

func printInfo() {
	// Print information about how to use this program
}

func checkExit(s string) {
	// Exit program if user typed 'exit'
}

func restartProgram() {
	// Ask user to hit Enter key to restart program
	// User can quit program by typing 'exit'
}

func printOutput() {
	// Print available move positions and a chaseboard
}
