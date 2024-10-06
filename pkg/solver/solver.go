package solver

import (
	"math"
	"sync"

	. "github.com/tliddle1/wordle"
	"github.com/tliddle1/wordle/data"
)

type ThomasSolver struct {
	validTargets []string
	validGuesses []string
}

func NewThomasSolver() *ThomasSolver {
	solver := ThomasSolver{}
	solver.setData()
	return &solver
}

func (this *ThomasSolver) Debug() bool {
	return false
}

func (this *ThomasSolver) Guess(turnHistory []Turn) string {
	if len(turnHistory) == 0 {
		return "soare"
	}
	this.updateValidTargets(turnHistory)
	guess := this.maximizeExpectedInformation()
	return guess
}

func (this *ThomasSolver) Reset() {
	this.setData()
}

// private

func (this *ThomasSolver) setData() {
	this.validTargets = data.ValidTargets
	this.validGuesses = append(data.ValidGuesses, data.ValidTargets...)
}

func (this *ThomasSolver) updateValidTargets(turnHistory []Turn) {
	if len(turnHistory) == 0 {
		return
	}
	lastTurn := turnHistory[len(turnHistory)-1]
	var newTargets []string
	for _, target := range this.validTargets {
		if this.isValidTarget(target, lastTurn) {
			newTargets = append(newTargets, target)
		}
	}
	this.validTargets = newTargets
}

func (this *ThomasSolver) isValidTarget(word string, turn Turn) bool {
	return CheckGuess(word, turn.Guess) == turn.Pattern
}

func (this *ThomasSolver) maximizeExpectedInformation() string {
	if len(this.validTargets) <= 2 {
		return this.validTargets[0]
	}
	wg := sync.WaitGroup{}
	// Channel for the guess and what the expected value is
	wordPairChannel := make(chan guessExpectedValuePair, 2000)
	wordWithMinExpectedValue := make(chan string)
	go determineWordWithMaxExpectedInfo(wordPairChannel, wordWithMinExpectedValue)
	for _, word := range this.validGuesses {
		wg.Add(1)
		go this.calculateExpectedInfoAndSendToCompareChannel(wordPairChannel, word, &wg)
	}
	wg.Wait()
	close(wordPairChannel)
	return <-wordWithMinExpectedValue
}

func (this *ThomasSolver) calculateExpectedInfoAndSendToCompareChannel(out chan guessExpectedValuePair, word string, wg *sync.WaitGroup) {
	defer wg.Done()
	expectedInfo := this.calculateExpectedInfo(word)
	out <- guessExpectedValuePair{word, expectedInfo}
}

func (this *ThomasSolver) calculateExpectedInfo(word string) float64 {
	possiblePatterns := make(map[Pattern]int)
	for _, possibleTarget := range this.validTargets {
		possiblePattern := CheckGuess(possibleTarget, word)
		possiblePatterns[possiblePattern]++
	}

	expectedInfo := float64(0)
	for _, count := range possiblePatterns {
		probabilityOfPattern := float64(count) / float64(len(this.validTargets))
		expectedInfo += -probabilityOfPattern * math.Log2(probabilityOfPattern)
	}
	return expectedInfo
}

func determineWordWithMaxExpectedInfo(in chan guessExpectedValuePair, out chan string) {
	var maxWord string
	maxVal := float64(0)
	for wordValuePair := range in {
		// if it's a tie, use the first guess alphabetically
		if wordValuePair.expectedValue > maxVal || (wordValuePair.expectedValue == maxVal && wordValuePair.guess < maxWord) {
			maxVal = wordValuePair.expectedValue
			maxWord = wordValuePair.guess
		}
	}
	out <- maxWord
}

type guessExpectedValuePair struct {
	guess         string
	expectedValue float64
}
