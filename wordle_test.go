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

// TODO table test
func (this *WordleFixture) TestGreenGrayAndYellow() {
	targetWord := "snake"
	guess := "slain"
	this.So(CheckGuess(targetWord, guess), should.Equal, Pattern{Green, Gray, Green, Gray, Yellow})
}

func (this *WordleFixture) TestDoubleLetterInAnswer() {
	targetWord := "sheen"
	guess := "siren"
	this.So(CheckGuess(targetWord, guess), should.Equal, Pattern{Green, Gray, Gray, Green, Green})
}

func (this *WordleFixture) TestTwoYellowsSameLetter() {
	targetWord := "sheen"
	guess := "elate"
	this.So(CheckGuess(targetWord, guess), should.Equal, Pattern{Yellow, Gray, Gray, Gray, Yellow})
}

func (this *WordleFixture) TestTwoLettersFirstYellow() {
	targetWord := "messy"
	guess := "sheen"
	this.So(CheckGuess(targetWord, guess), should.Equal, Pattern{Yellow, Gray, Yellow, Gray, Gray})
}

func (this *WordleFixture) TestLotsOfRepeatedLetters() {
	targetWord := "freer"
	guess := "error"
	this.So(CheckGuess(targetWord, guess), should.Equal, Pattern{Yellow, Green, Gray, Gray, Green})
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
