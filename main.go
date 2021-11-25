package main

import (
	"fmt"
)

var wordstoguess string
var hiddenWord = []rune{}


func initWordToGuess(word string) {
	wordstoguess = word
}

func initHiddenWord() {
	for range wordstoguess {
		hiddenWord = append(hiddenWord, '_')
	}
}

func Contains(letter rune) {
	for i, char := range wordstoguess {
		if char == letter {
			hiddenWord[i] = char
		}
	}
}

func Printword() {
	for _, char := range hiddenWord {
		fmt.Print(string(char))
		fmt.Print(" ")
	}

	fmt.Println()
}

func main() {
	initWordToGuess("KheirEddinE")
	initHiddenWord()
	Contains('h')
	Printword()
	Contains('K')
	Printword()
	Contains('E')
	Printword()
}