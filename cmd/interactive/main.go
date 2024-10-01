package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/tliddle1/wordle"
	"github.com/tliddle1/wordle/data"
)

func main() {
	evaluator := wordle.NewEvaluator()
	avgNumGuesses, err := evaluator.EvaluateSolver(interactiveSolver{append(data.ValidTargets, data.ValidGuesses...)})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(avgNumGuesses)
}

type interactiveSolver struct {
	validGuesses []string
}

func (this interactiveSolver) Debug() bool {
	return false
}

func (this interactiveSolver) Guess(guesses []string, patterns []wordle.Pattern) string {
	if len(guesses) != len(patterns) {
		panic("guesses and patterns have different length")
	}
	if len(guesses) > 0 {
		wordle.PrintPattern(patterns[len(patterns)-1], guesses[len(guesses)-1])
	}
	reader := bufio.NewReader(os.Stdin)
	validGuess := false
	var guess string
	for !validGuess {
		fmt.Print("Enter your guess: ")
		input, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(input)
		if !slices.Contains(this.validGuesses, guess) {
			fmt.Println("Invalid guess, try again.")
		} else {
			validGuess = true
		}
	}
	return guess
}

func (this interactiveSolver) Reset() {}
