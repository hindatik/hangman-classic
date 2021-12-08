package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var words = make([]string, 0, 2)

// va lire le fichier en argument 

func Load(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	word := ""

	for _, char := range string(f) {
		if char == '\n' {
			words = append(words, word)
			word = ""
		} else {
			word += string(char)
		}
	}
}

func PickWord() string {

	// Pour Pouvoir Choisir un mot random
	
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(words))
	return words[i]
}
