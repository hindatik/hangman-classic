package main

import (
	"fmt"
	"os"
)
func main() {
	
	Load(os.Args[1])

	g, err := New(10, PickWord())

	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}
	
	DrawWelcome()
	guess := ""
	for {
		Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l

		g.MakeAGuess(guess)
	}
}
