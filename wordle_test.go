package wordle

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestWordleFixture(t *testing.T) {
	gunit.Run(new(WordleFixture), t)
}

type WordleFixture struct {
	*gunit.Fixture
	Evaluator *Evaluator
}

func (this *WordleFixture) Setup() {
	this.Evaluator = NewEvaluator()
}

func (this *WordleFixture) TestEvaluatorOneGuess() {
	avgNumGuesses, err := this.Evaluator.EvaluateSolver(NewDummySolverOneGuess())
	this.So(err, should.Wrap, ErrLostGame)
	this.So(avgNumGuesses, should.Equal, -1)
}

func (this *WordleFixture) NewDummySolverInvalidGuess() {
	avgNumGuesses, err := this.Evaluator.EvaluateSolver(NewDummySolverInvalidGuess())
	this.So(err, should.Equal, ErrInvalidGuess)
	this.So(avgNumGuesses, should.Equal, -1)
}

func TestIsValidTarget(t *testing.T) {
	tests := []struct {
		name     string
		target   string
		guess    string
		expected Pattern
	}{
		{name: "green, gray, and yellow", target: "snake", guess: "slain", expected: Pattern{Green, Gray, Green, Gray, Yellow}},
		{name: "double letter in answer", target: "sheen", guess: "siren", expected: Pattern{Green, Gray, Gray, Green, Green}},
		{name: "two yellows same letter", target: "sheen", guess: "elate", expected: Pattern{Yellow, Gray, Gray, Gray, Yellow}},
		{name: "two letters one yellow", target: "messy", guess: "sheen", expected: Pattern{Yellow, Gray, Yellow, Gray, Gray}},
		{name: "lots of repeated letters", target: "freer", guess: "error", expected: Pattern{Yellow, Green, Gray, Gray, Green}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckGuess(tt.target, tt.guess)
			if result != tt.expected {
				t.Errorf("isValidTarget returned %v when %v was expected for target: %s and guess: %s", result, tt.expected, tt.target, tt.guess)
			}
		})
	}
}

////////////////////////////////////////////////////////////////////////////////

type DummySolverOneGuess struct{}

func NewDummySolverOneGuess() Solver {
	return &DummySolverOneGuess{}
}

func (this DummySolverOneGuess) Debug() bool {
	return false
}

func (this DummySolverOneGuess) Guess(turnHistory []Turn) string {
	return "salet"
}

func (this DummySolverOneGuess) Reset() {}

////////////////////////////////////////////////////////////////////////////////

type DummySolverInvalidGuess struct{}

func NewDummySolverInvalidGuess() Solver {
	return &DummySolverInvalidGuess{}
}

func (this DummySolverInvalidGuess) Debug() bool {
	return false
}

func (this DummySolverInvalidGuess) Guess(turnHistory []Turn) string {
	return "sssss"
}

func (this DummySolverInvalidGuess) Reset() {}
