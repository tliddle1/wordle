package wordle

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/tliddle1/wordle/data"
	"github.com/tliddle1/wordle/pkg/set"
)

type Solver interface {
	// Debug returns true if the solver is in debug mode
	Debug() bool
	// Guess returns the next guess from the solver given the current turn history
	Guess(turnHistory []Turn) string
	// Reset will reset the original state of the solver between games
	Reset()
}

var (
	ErrInvalidGuess       = errors.New("invalid guess")
	ErrInvalidLengthGuess = errors.New("guess is not 5 letters")
	ErrLostGame           = errors.New("a game took longer than the maximum number of guesses")
	CorrectPattern        = Pattern{Green, Green, Green, Green, Green}
	grayPattern           = Pattern{Gray, Gray, Gray, Gray, Gray}
)

const (
	WordLength    = 5
	MaxNumGuesses = 6
)

const (
	Gray LetterColor = iota
	Yellow
	Green
)

type LetterColor int

// Pattern is the clue for a wordle guess
type Pattern [5]LetterColor

// Turn is a guess with its respective pattern
type Turn struct {
	Guess   string  // the word that was guessed
	Pattern Pattern // the pattern returned by the wordle game for that guess
}

type Evaluator struct {
	validTargetSlice []string
	validGuessSet    set.Set[string]
}

func NewEvaluator() *Evaluator {

	evaluator := Evaluator{
		validTargetSlice: make([]string, len(data.ValidTargets)),
		validGuessSet:    set.Set[string]{},
	}
	copy(evaluator.validTargetSlice, data.ValidTargets)
	for _, guess := range data.ValidTargets {
		evaluator.validGuessSet.Add(guess)
	}
	for _, guess := range data.ValidGuesses {
		evaluator.validGuessSet.Add(guess)
	}
	rand.Shuffle(len(evaluator.validTargetSlice), func(i, j int) {
		evaluator.validTargetSlice[i], evaluator.validTargetSlice[j] = evaluator.validTargetSlice[j], evaluator.validTargetSlice[i]
	})
	return &evaluator
}

// EvaluateSolver will return the average number of guesses that a solver needs to solve all wordles
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
		totalGuesses += numGuesses
		mostGuesses = max(numGuesses, mostGuesses)
		if debug {
			break
		}
	}
	fmt.Printf("Longest game took %d guesses.\n", mostGuesses)
	return float32(totalGuesses) / float32(len(this.validTargetSlice)), nil
}

// PlayGame will simulate a single game of wordle
func (this *Evaluator) PlayGame(target string, solver Solver) (int, error) {
	debug := solver.Debug()
	var turnHistory []Turn

	for i := 1; i <= MaxNumGuesses; i++ {
		guess := solver.Guess(turnHistory)
		if len(guess) != WordLength {
			return -1, fmt.Errorf("%w: \"%s\"", ErrInvalidLengthGuess, guess)
		}
		if !this.validGuessSet.Contains(guess) {
			return -1, fmt.Errorf("%w: \"%s\"", ErrInvalidGuess, guess)
		}

		pattern := CheckGuess(target, guess)
		if pattern == CorrectPattern {
			if debug {
				PrintPattern(CorrectPattern, target)
				fmt.Println(i, "guesses")
			}
			return i, nil
		}
		if debug {
			PrintPattern(pattern, guess)
		}
		turnHistory = append(turnHistory, Turn{guess, pattern})
	}
	if debug {
		fmt.Printf("The word was: %s\n", target)
	}
	return MaxNumGuesses, fmt.Errorf("%w: %s", ErrLostGame, target)
}

// CheckGuess will return the pattern of a guess for a particular target
func CheckGuess(target, guess string) Pattern {
	return checkGuess([]byte(target), []byte(guess))
}

func checkGuess(target, guess []byte) Pattern {
	used := make([]bool, WordLength)
	pattern := grayPattern

	// First pass: Check for exact matches (Green patterns)
	for i := 0; i < WordLength; i++ {
		if target[i] == guess[i] {
			pattern[i] = Green
			used[i] = true // Mark this position as used
		}
	}

	// Second pass: Check for partial matches (Yellow patterns)
	for i := 0; i < WordLength; i++ {
		if pattern[i] == Gray {
			for j := 0; j < WordLength; j++ {
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

// PrintPattern will print the guess using the colors from the pattern for each letter
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
