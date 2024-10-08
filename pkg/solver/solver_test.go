package solver

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	. "github.com/tliddle1/wordle"
)

func TestSolverFixture(t *testing.T) {
	gunit.Run(new(SolverFixture), t)
}

type SolverFixture struct {
	*gunit.Fixture
	Solver    *ThomasSolver
	Evaluator *Evaluator
}

func (this *SolverFixture) Setup() {
	this.Solver = NewThomasSolver()
	this.Evaluator = NewEvaluator()
}

func TestIsValidTarget(t *testing.T) {
	solver := NewThomasSolver()
	tests := []struct {
		name       string
		targetWord string
		turn       Turn
		expected   bool
	}{
		{name: "some green, yellow, and gray", targetWord: "snake", turn: Turn{Guess: "slain", Pattern: Pattern{Green, Gray, Green, Gray, Yellow}}, expected: true},
		{name: "duplicate letter 1 yellow 1 gray", targetWord: "stare", turn: Turn{Guess: "steer", Pattern: Pattern{Green, Green, Yellow, Gray, Yellow}}, expected: true},
		{name: "duplicate letter in target", targetWord: "sheen", turn: Turn{Guess: "siren", Pattern: Pattern{Green, Gray, Gray, Green, Green}}, expected: true},
		{name: "duplicate letter both yellow", targetWord: "sheen", turn: Turn{Guess: "elate", Pattern: Pattern{Yellow, Gray, Gray, Gray, Yellow}}, expected: true},
		{name: "duplicate letter in target and guess", targetWord: "messy", turn: Turn{Guess: "sheen", Pattern: Pattern{Yellow, Gray, Yellow, Gray, Gray}}, expected: true},
		{name: "triple letter target", targetWord: "error", turn: Turn{Guess: "revel", Pattern: Pattern{Yellow, Yellow, Gray, Gray, Gray}}, expected: true},
		{name: "triple letter guess", targetWord: "freer", turn: Turn{Guess: "error", Pattern: Pattern{Yellow, Green, Gray, Gray, Green}}, expected: true},
		{name: "duplicate letter both yellow false", targetWord: "stare", turn: Turn{Guess: "bleed", Pattern: Pattern{Gray, Gray, Yellow, Yellow, Gray}}, expected: false},
		{name: "only 1 letter yellow falsea", targetWord: "cater", turn: Turn{Guess: "blech", Pattern: Pattern{Gray, Gray, Yellow, Gray, Gray}}, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solver.isValidTarget(tt.targetWord, tt.turn)
			if result != tt.expected {
				t.Errorf("isValidTarget returned %v when %v was expected for target: %s and guess: %s", result, tt.expected, tt.targetWord, tt.turn.Guess)
			}
		})
	}
}

func (this *SolverFixture) TestCalculateExpectedInformationSmart() {
	word := "smart"
	expectedInformation := this.Solver.calculateExpectedInfo(word)
	this.So(expectedInformation, should.AlmostEqual, 5.24325852268321)
}

func (this *SolverFixture) TestCalculateExpectedInformationRaise() {
	word := "raise"
	expectedInformation := this.Solver.calculateExpectedInfo(word)
	this.So(expectedInformation, should.AlmostEqual, 5.87830295649317)
}

func (this *SolverFixture) TestCalculateExpectedInformationSoare() {
	word := "soare"
	expectedInformation := this.Solver.calculateExpectedInfo(word)
	this.So(expectedInformation, should.AlmostEqual, 5.8852027442927)
}

func (this *SolverFixture) TestSingleGame() {
	numGuesses, err := this.Evaluator.PlayGame("angry", this.Solver)
	this.So(err, should.BeNil)
	this.So(numGuesses, should.BeLessThanOrEqualTo, MaxNumGuesses)
}

func (this *SolverFixture) TestUpdateValidTargetsNoOp() {
	preUpdateLength := len(this.Solver.validTargets)
	this.Solver.updateValidTargets([]Turn{})
	this.So(this.Solver.validTargets, should.HaveLength, preUpdateLength)
}

func (this *SolverFixture) TestReset() {
	preUpdateLength := len(this.Solver.validTargets)
	this.Solver.updateValidTargets([]Turn{{Guess: "soare", Pattern: Pattern{Gray, Gray, Gray, Gray, Gray}}})
	this.So(len(this.Solver.validTargets), should.BeLessThan, preUpdateLength)
	this.Solver.Reset()
	this.So(this.Solver.validTargets, should.HaveLength, preUpdateLength)
}
