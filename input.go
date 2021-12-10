package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
/// va lire les inputs
func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("What is your letter? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		guess = strings.TrimSpace(guess)

	
		valid = true
	}
	return guess, nil
}
