package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"

	"github.com/tliddle1/wordle"
	"github.com/tliddle1/wordle/data"
)

var validGuesses = append(data.ValidTargets, data.ValidGuesses...)

func main() {
	target := data.ValidTargets[rand.Intn(len(data.ValidTargets))]
	scanner := bufio.NewScanner(os.Stdin)
	won := false
	for range wordle.MaxNumGuesses {
		guess := askForGuess(scanner)
		pattern := wordle.CheckGuess(target, guess)
		wordle.PrintPattern(pattern, guess)
		if pattern == wordle.CorrectPattern {
			won = true
			break
		}
	}
	if won {
		fmt.Println("You won!")
	} else {
		fmt.Printf("Sorry, you lost. The answer was %s.\n", target)
	}
}

func askForGuess(scanner *bufio.Scanner) (guess string) {
	validGuess := false
	for !validGuess {
		fmt.Print("Enter your guess: ")
		scanner.Scan()
		guess = scanner.Text()
		if !slices.Contains(validGuesses, guess) {
			fmt.Println("Invalid guess, try again.")
		} else {
			validGuess = true
		}
	}
	return guess
}
