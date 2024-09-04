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
	Solver *ThomasSolver
}

func (this *SolverFixture) Setup() {
	this.Solver = NewThomasSolver()
}

func (this *SolverFixture) Test() {
	targetWord := "snake"
	guess := "slain"
	var clue Clue
	clue = [5]LetterColor{Green, Gray, Green, Gray, Yellow}
	valid := this.Solver.isValidWord(targetWord, guess, clue)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test2() {
	targetWord := "stare"
	guess := "steer"
	var clue Clue
	clue = [5]LetterColor{Green, Green, Yellow, Gray, Yellow}
	valid := this.Solver.isValidWord(targetWord, guess, clue)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test3() {
	targetWord := "sheen"
	guess := "siren"
	var clue Clue
	clue = [5]LetterColor{Green, Gray, Gray, Green, Green}
	valid := this.Solver.isValidWord(targetWord, guess, clue)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test4() {
	targetWord := "sheen"
	guess := "elate"
	var clue Clue
	clue = [5]LetterColor{Yellow, Gray, Gray, Gray, Yellow}
	valid := this.Solver.isValidWord(targetWord, guess, clue)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test5() {
	targetWord := "messy"
	guess := "sheen"
	var clue Clue
	clue = [5]LetterColor{Yellow, Gray, Yellow, Gray, Gray}
	valid := this.Solver.isValidWord(targetWord, guess, clue)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test6() {
	targetWord := "stare"
	guess := "bleed"
	var clue Clue
	clue = [5]LetterColor{Gray, Gray, Yellow, Yellow, Gray}
	valid := this.Solver.isValidWord(targetWord, guess, clue)
	this.So(valid, should.BeFalse)
}

func (this *SolverFixture) Test7() {
	targetWord := "elate"
	guess := "bleed"
	var clue Clue
	clue = [5]LetterColor{Gray, Green, Yellow, Yellow, Gray}
	valid := this.Solver.isValidWord(targetWord, guess, clue)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test8() {
	targetWord := "error"
	guess := "revel"
	var clue Clue
	clue = [5]LetterColor{Yellow, Yellow, Gray, Gray, Gray}
	valid := this.Solver.isValidWord(targetWord, guess, clue)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test9() {
	targetWord := "freer"
	guess := "error"
	var clue Clue
	clue = [5]LetterColor{Yellow, Green, Gray, Gray, Green}
	valid := this.Solver.isValidWord(targetWord, guess, clue)
	this.So(valid, should.BeTrue)
}
