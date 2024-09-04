package solver

import (
	. "github.com/tliddle1/wordle"
	"github.com/tliddle1/wordle/data"
)

type ThomasSolver struct {
	validTargets []string
	validGuesses []string
}

func NewThomasSolver() *ThomasSolver {
	solver := ThomasSolver{
		validTargets: make([]string, len(data.ValidTargets)),
		validGuesses: make([]string, len(data.ValidGuesses)),
	}
	copy(solver.validTargets, data.ValidTargets)
	copy(solver.validGuesses, data.ValidGuesses)
	return &solver
}

func (this *ThomasSolver) Debug() bool {
	return false
}

func (this *ThomasSolver) Guess(guesses []string, clues []Clue) string {
	if len(clues) == 0 {
		return "train"
	}
	this.updateValidTargets(clues[len(clues)-1], guesses[len(guesses)-1])
	return this.validTargets[0]
}

func (this *ThomasSolver) Reset() {
	this.validTargets = data.ValidTargets
	this.validGuesses = data.ValidGuesses
}

func (this *ThomasSolver) updateValidTargets(clue Clue, guess string) {
	var newTargets []string
	for _, target := range this.validTargets {
		if this.isValidWord(target, guess, clue) {
			newTargets = append(newTargets, target)
		}
	}
	this.validTargets = newTargets
}

func (this *ThomasSolver) isValidWord(word, guess string, clue Clue) bool {
	letterCount := make(map[rune]int)
	letterCountDebug := make(map[string]int)

	for i, letter := range guess {
		if clue[i] == Green || clue[i] == Yellow {
			letterCountDebug[string(letter)]++
			letterCount[letter]++
		}
	}

	for i := range 5 {
		switch clue[i] {
		case Green:
			if word[i] != guess[i] {
				return false
			}
		case Yellow:
			if word[i] == guess[i] {
				return false
			}
		case Gray:
			if word[i] == guess[i] {
				return false
			}
		}
	}
	for _, letter := range word {
		letterCount[letter]--
	}
	for _, count := range letterCount {
		if count > 0 {
			return false
		}
	}
	return true
}
