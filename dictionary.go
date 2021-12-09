package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var words = []string{}

// va lire le fichier en argument 

func Load(filename string) bool{
	if strings.TrimSpace(filename) == "" {
		return false
	}
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
	return true
}

func PickWord() string {

	// Pour Pouvoir Choisir un mot random

	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(words))
	return words[i]
}
