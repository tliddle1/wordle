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
	this.So(err, should.BeNil)
	this.So(avgNumGuesses, should.Equal, maxNumGuesses+1)
}

func (this *WordleFixture) NewDummySolverInvalidGuess() {
	avgNumGuesses, err := this.Evaluator.EvaluateSolver(NewDummySolverInvalidGuess())
	this.So(err, should.Equal, ErrInvalidGuess)
	this.So(avgNumGuesses, should.Equal, -1)
}

func (this *WordleFixture) TestGreenGrayAndYellow() {
	targetWord := []byte{'s', 'n', 'a', 'k', 'e'}
	guess := []byte{'s', 'l', 'a', 'i', 'n'}
	var expectedPattern Pattern
	expectedPattern = [5]LetterColor{Green, Gray, Green, Gray, Yellow}
	pattern := checkGuess(targetWord, guess)
	this.So(pattern, should.Equal, expectedPattern)
}

func (this *WordleFixture) TestDoubleLetterInAnswer() {
	targetWord := []byte{'s', 'h', 'e', 'e', 'n'}
	guess := []byte{'s', 'i', 'r', 'e', 'n'}
	var expectedPattern Pattern
	expectedPattern = [5]LetterColor{Green, Gray, Gray, Green, Green}
	pattern := checkGuess(targetWord, guess)
	this.So(pattern, should.Equal, expectedPattern)
}

func (this *WordleFixture) TestTwoYellowsSameLetter() {
	targetWord := []byte{'s', 'h', 'e', 'e', 'n'}
	guess := []byte{'e', 'l', 'a', 't', 'e'}
	var expectedPattern Pattern
	expectedPattern = [5]LetterColor{Yellow, Gray, Gray, Gray, Yellow}
	pattern := checkGuess(targetWord, guess)
	this.So(pattern, should.Equal, expectedPattern)
}

func (this *WordleFixture) TestTwoLettersFirstYellow() {
	targetWord := []byte{'m', 'e', 's', 's', 'y'}
	guess := []byte{'s', 'h', 'e', 'e', 'n'}
	var expectedPattern Pattern
	expectedPattern = [5]LetterColor{Yellow, Gray, Yellow, Gray, Gray}
	pattern := checkGuess(targetWord, guess)
	this.So(pattern, should.Equal, expectedPattern)
}

func (this *WordleFixture) Test2() {
	targetWord := []byte{'f', 'r', 'e', 'e', 'r'}
	guess := []byte{'e', 'r', 'r', 'o', 'r'}
	var expectedPattern Pattern
	expectedPattern = [5]LetterColor{Yellow, Green, Gray, Gray, Green}
	pattern := checkGuess(targetWord, guess)
	this.So(pattern, should.Equal, expectedPattern)
}

////////////////////////////////////////////////////////////////////////////////

type DummySolverOneGuess struct{}

func (this DummySolverOneGuess) Reset() {}

func NewDummySolverOneGuess() Solver {
	return &DummySolverOneGuess{}
}

func (this DummySolverOneGuess) Debug() bool {
	return false
}

func (this DummySolverOneGuess) Guess(guesses []string, patterns []Pattern) string {
	return "salet"
}

type DummySolverInvalidGuess struct{}

func NewDummySolverInvalidGuess() Solver {
	return &DummySolverInvalidGuess{}
}

func (this DummySolverInvalidGuess) Debug() bool {
	return false
}

func (this DummySolverInvalidGuess) Guess(guesses []string, patterns []Pattern) string {
	return "sssss"
}

func (this DummySolverInvalidGuess) Reset() {}
