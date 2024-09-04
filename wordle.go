package wordle

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"slices"

	"github.com/tliddle1/wordle/data"
)

type Solver interface {
	Debug() bool
	Guess(guesses []string, clues []Clue) string
	Reset()
}

type LetterColor int

var (
	correct               = [5]LetterColor{Green, Green, Green, Green, Green}
	blankClue             = [5]LetterColor{unknown, unknown, unknown, unknown, unknown}
	ErrInvalidGuess       = errors.New("invalid guess")
	ErrInvalidLengthGuess = errors.New("guess is not 5 letters")
)

const (
	wordLength    = 5
	maxNumGuesses = 3000
)

const (
	Gray LetterColor = iota
	Yellow
	Green
	unknown
)

type Clue [5]LetterColor

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
	for _, targetString := range this.validTargetSlice {
		solver.Reset()
		if debug {
			fmt.Println("target:", targetString)
		}
		numGuesses, err := this.playGame(targetString, solver)
		if err != nil {
			return -1, err
		}
		totalGuesses += numGuesses
		if debug {
			break
		}
	}
	return float32(totalGuesses) / float32(len(this.validTargetSlice)), nil
}

func (this *Evaluator) playGame(targetString string, solver Solver) (int, error) {
	debug := solver.Debug()
	var guesses []string
	var clues []Clue

	for i := 1; i <= maxNumGuesses; i++ {
		guessString := solver.Guess(guesses, clues)
		if len(guessString) != wordLength {
			return -1, fmt.Errorf("%w: %s", ErrInvalidLengthGuess, guessString)
		}
		if !this.validGuessMap[guessString] && !this.validTargetMap[guessString] {
			return -1, fmt.Errorf("%w: %s", ErrInvalidGuess, guessString)
		}

		clue := checkGuess([]byte(targetString), []byte(guessString))
		if clue == correct {
			if debug {
				PrintClue(correct, targetString)
				fmt.Println(i, "guesses")
			}
			return i, nil
		}
		if debug {
			PrintClue(clue, guessString)
		}
		guesses = append(guesses, guessString)
		clues = append(clues, clue)
	}
	if debug {
		fmt.Printf("The word was: %s\n", targetString)
	}
	return maxNumGuesses + 1, nil
}

func checkGuess(target, guess []byte) (clue Clue) {
	guess = bytes.ToLower(guess)
	clue = blankClue
	for i := range 5 {
		if target[i] == guess[i] {
			clue[i] = Green
			target[i] = 0
		}
	}
	for i := range 5 {
		if clue[i] == unknown && slices.Contains(target, guess[i]) {
			clue[i] = Yellow
			idx := bytes.IndexByte(target, guess[i])
			target[idx] = 0
		}
	}
	for i := range 5 {
		if clue[i] == unknown {
			clue[i] = Gray
		}
	}
	return clue
}

func PrintClue(clue Clue, guess string) {
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	colorizedClue := ""
	for i := range 5 {
		if clue[i] == Green {
			colorizedClue += green + string(guess[i]) + reset
		} else if clue[i] == Yellow {
			colorizedClue += yellow + string(guess[i]) + reset
		} else {
			colorizedClue += string(guess[i])
		}
	}
	fmt.Println(colorizedClue)
}
