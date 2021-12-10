package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)
///affiche le nom du jeu 

func DrawWelcome() {
	fmt.Println(`
	|\     /|(  ___  )( (    /|(  ____ \(       )(  ___  )( (    /|
	| )   ( || (   ) ||  \  ( || (    \/| () () || (   ) ||  \  ( |
	| (___) || (___) ||   \ | || |      | || || || (___) ||   \ | |
	|  ___  ||  ___  || (\ \) || | ____ | |(_)| ||  ___  || (\ \) |
	| (   ) || (   ) || | \   || | \_  )| |   | || (   ) || | \   |
	| )   ( || )   ( || )  \  || (___) || )   ( || )   ( || )  \  |
	|/     \||/     \||/    )_)(_______)|/     \||/     \||/    )_)
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
	fmt.Println()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

///ca nous sert a afficher le petit pendu
func drawTurns(l int) {
	file, err := os.Open("hangman.txt")
	errfunc(err)
    fileScanner := bufio.NewScanner(file)
    i := 0
    var debut int
    var fin int
    switch l {
    case 9:
        debut = 0
        fin = 7
    case 8:
        debut = 8
        fin = 15
    case 7:
        debut = 16
        fin = 23
    case 6:
        debut = 24
        fin = 31
    case 5:
        debut = 32
        fin = 39
    case 4:
        debut = 40
        fin = 47
    case 3:
        debut = 48
        fin = 55
    case 2:
        debut = 56
        fin = 63
    case 1:
        debut = 64
        fin = 71
    case 0:
        debut = 72
        fin = 79
    }
    for fileScanner.Scan() {
        if i >= debut && i <= fin {
            println(fileScanner.Text())
        }
        i++
    }
}

func drawLetters(g []string) {
	for _, c := range g {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}

func errfunc(err error) {
    if err != nil {
        log.Fatalf("Error when opening file: %s", err)
    }
}
//cette fonction nous permet de suivre l'état de la partie 
func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	fmt.Print("Turn Left: ")
	fmt.Println(g.TurnsLeft)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word", guess)
	case "lost":
		fmt.Print("You lost :(! The word was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("YOU WON! The word was: ")
		drawLetters(g.Letters)
	}
	fmt.Println()
}
