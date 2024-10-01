package wordle

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/tliddle1/wordle/data"
)

type Solver interface {
	Debug() bool
	Guess(guesses []string, patterns []Pattern) string
	Reset()
}

type LetterColor int

var (
	correct               = [5]LetterColor{Green, Green, Green, Green, Green}
	grayPattern           = [5]LetterColor{Gray, Gray, Gray, Gray, Gray}
	ErrInvalidGuess       = errors.New("invalid guess")
	ErrInvalidLengthGuess = errors.New("guess is not 5 letters")
	ErrLostGame           = errors.New("a game took longer than the maximum number of guesses")
)

const (
	wordLength    = 5
	maxNumGuesses = 6
)

const (
	Gray LetterColor = iota
	Yellow
	Green
)

type Pattern [5]LetterColor

type Evaluator struct {
	validTargetSlice []string
	validTargetMap   map[string]bool
	validGuessMap    map[string]bool
}

func NewEvaluator() *Evaluator {

	evaluator := Evaluator{
		validTargetSlice: make([]string, len(data.ValidTargets)),
		validTargetMap:   map[string]bool{},
		validGuessMap:    map[string]bool{},
	}
	copy(evaluator.validTargetSlice, data.ValidTargets)
	for _, target := range evaluator.validTargetSlice {
		evaluator.validTargetMap[target] = true
	}
	for _, guess := range data.ValidGuesses {
		evaluator.validGuessMap[guess] = true
	}
	rand.Shuffle(len(evaluator.validTargetSlice), func(i, j int) {
		evaluator.validTargetSlice[i], evaluator.validTargetSlice[j] = evaluator.validTargetSlice[j], evaluator.validTargetSlice[i]
	})
	return &evaluator
}

func (this *Evaluator) EvaluateSolver(solver Solver) (float32, error) {
	debug := solver.Debug()
	var totalGuesses int
	i := 0
	mostGuesses := 0
	for _, targetString := range this.validTargetSlice {
		i++
		if i%50 == 0 {
			fmt.Printf("%d/%d completed\n", i, len(this.validTargetSlice))
		}
		solver.Reset()
		if debug {
			fmt.Println("target:", targetString)
		}
		numGuesses, err := this.PlayGame(targetString, solver)
		if err != nil {
			return -1, err
		}
		if numGuesses > maxNumGuesses {
			return -1, ErrLostGame
		}
		totalGuesses += numGuesses
		mostGuesses = max(numGuesses, mostGuesses)
		if debug {
			break
		}
	}
	fmt.Printf("Longest game took %d guesses.\n", mostGuesses)
	return float32(totalGuesses) / float32(len(this.validTargetSlice)), nil
}

func (this *Evaluator) PlayGame(target string, solver Solver) (int, error) {
	debug := solver.Debug()
	var guesses []string
	var patterns []Pattern

	for i := 1; i <= maxNumGuesses; i++ {
		guess := solver.Guess(guesses, patterns)
		if len(guess) != wordLength {
			return -1, fmt.Errorf("%w: \"%s\"", ErrInvalidLengthGuess, guess)
		}
		if !this.validGuessMap[guess] && !this.validTargetMap[guess] {
			return -1, fmt.Errorf("%w: \"%s\"", ErrInvalidGuess, guess)
		}

		pattern := CheckGuess(target, guess)
		if pattern == correct {
			if debug {
				PrintPattern(correct, target)
				fmt.Println(i, "guesses")
			}
			return i, nil
		}
		if debug {
			PrintPattern(pattern, guess)
		}
		guesses = append(guesses, guess)
		patterns = append(patterns, pattern)
	}
	if debug {
		fmt.Printf("The word was: %s\n", target)
	}
	return maxNumGuesses + 1, nil
}

func CheckGuess(target, guess string) (pattern Pattern) {
	return checkGuess([]byte(target), []byte(guess))
}

func checkGuess(target, guess []byte) (pattern Pattern) {
	used := make([]bool, wordLength)
	pattern = grayPattern

	// First pass: Check for exact matches (Green patterns)
	for i := 0; i < wordLength; i++ {
		if target[i] == guess[i] {
			pattern[i] = Green
			used[i] = true // Mark this position as used
		}
	}

	// Second pass: Check for partial matches (Yellow patterns)
	for i := 0; i < wordLength; i++ {
		if pattern[i] == Gray {
			for j := 0; j < wordLength; j++ {
				if !used[j] && target[j] == guess[i] {
					pattern[i] = Yellow
					used[j] = true // Mark this position as used
					break
				}
			}
		}
	}

	return pattern
}

func PrintPattern(pattern Pattern, guess string) {
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	colorizedPattern := ""
	for i := range 5 {
		if pattern[i] == Green {
			colorizedPattern += green + string(guess[i]) + reset
		} else if pattern[i] == Yellow {
			colorizedPattern += yellow + string(guess[i]) + reset
		} else {
			colorizedPattern += string(guess[i])
		}
	}
	fmt.Println(colorizedPattern)
}
