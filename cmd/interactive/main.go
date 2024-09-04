package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tliddle1/wordle"
)

func main() {
	evaluator := wordle.NewEvaluator()
	avgNumGuesses, err := evaluator.EvaluateSolver(interactiveSolver{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(avgNumGuesses)
}

type interactiveSolver struct{}

func (this interactiveSolver) Debug() bool {
	return false
}

func (this interactiveSolver) Guess(guesses []string, clues []wordle.Clue) string {
	if len(guesses) != len(clues) {
		panic("guesses and clues have different length")
	}
	if len(guesses) > 0 {
		wordle.PrintClue(clues[len(clues)-1], guesses[len(guesses)-1])
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your guess: ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (this interactiveSolver) Reset() {}
