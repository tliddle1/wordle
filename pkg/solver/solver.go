package solver

import (
	"math"
	"sync"

	. "github.com/tliddle1/wordle"
	"github.com/tliddle1/wordle/data"
)

type ThomasSolver struct {
	validTargetsLeft []string
	validGuesses     []string
}

func NewThomasSolver() *ThomasSolver {
	solver := ThomasSolver{
		validTargetsLeft: make([]string, len(data.ValidTargets)),
		validGuesses:     make([]string, len(data.ValidGuesses)+len(data.ValidTargets)),
	}
	copy(solver.validTargetsLeft, data.ValidTargets)
	solver.validGuesses = append(data.ValidGuesses, data.ValidTargets...)
	return &solver
}

func (this *ThomasSolver) Debug() bool {
	return false
}

func (this *ThomasSolver) Guess(turnHistory []Turn) string {
	if len(turnHistory) == 0 {
		return "soare"
	}
	this.updateValidTargets(turnHistory[len(turnHistory)-1])
	if len(this.validTargetsLeft) == 1 {
		return this.validTargetsLeft[0]
	}
	guess := this.maximizeExpectedInformation()
	return guess
}

func (this *ThomasSolver) Reset() {
	this.validTargetsLeft = data.ValidTargets
	this.validGuesses = data.ValidGuesses
}

// private

func (this *ThomasSolver) updateValidTargets(turn Turn) {
	var newTargets []string
	for _, target := range this.validTargetsLeft {
		if this.isValidTarget(target, turn.Guess, turn.Pattern) {
			newTargets = append(newTargets, target)
		}
	}
	this.validTargetsLeft = newTargets
}

func (this *ThomasSolver) isValidTarget(word, guess string, pattern Pattern) bool {
	return CheckGuess(word, guess) == pattern
}

func (this *ThomasSolver) maximizeExpectedInformation() string {
	wg := sync.WaitGroup{}
	// Channel for the guess and what the expected value is
	wordPairChannel := make(chan wordExpectedValuePair, 40)
	wordWithMinExpectedValue := make(chan string)
	go determineWordWithMaxExpectedInfo(wordPairChannel, wordWithMinExpectedValue)
	// Try every possible
	for _, word := range this.validTargetsLeft {
		wg.Add(1)
		go this.calculateExpectedInfoAndSendToCompareChannel(wordPairChannel, word, &wg)
	}
	for _, word := range this.validGuesses {
		wg.Add(1)
		go this.calculateExpectedInfoAndSendToCompareChannel(wordPairChannel, word, &wg)
	}
	wg.Wait()
	close(wordPairChannel)
	return <-wordWithMinExpectedValue
}

func (this *ThomasSolver) calculateExpectedInfoAndSendToCompareChannel(out chan wordExpectedValuePair, word string, wg *sync.WaitGroup) {
	defer wg.Done()
	expectedInfo := this.calculateExpectedInfo(word)
	out <- wordExpectedValuePair{word, expectedInfo}
}

func (this *ThomasSolver) calculateExpectedInfo(word string) float64 {
	possiblePatterns := make(map[Pattern]int)
	for _, possibleTarget := range this.validTargetsLeft {
		pattern := CheckGuess(possibleTarget, word)
		possiblePatterns[pattern]++
	}

	expectedInfo := float64(0)
	for _, count := range possiblePatterns {
		probabilityOfPattern := float64(count) / float64(len(this.validTargetsLeft))
		expectedInfo += probabilityOfPattern * math.Log2(1/probabilityOfPattern)
	}
	return expectedInfo
}

func determineWordWithMaxExpectedInfo(in chan wordExpectedValuePair, out chan string) {
	var maxWord string
	maxVal := float64(0)
	for wordValuePair := range in {
		if wordValuePair.expectedValue > maxVal {
			maxVal = wordValuePair.expectedValue
			maxWord = wordValuePair.word
		}
	}
	out <- maxWord
}

type wordExpectedValuePair struct {
	word          string
	expectedValue float64
}
