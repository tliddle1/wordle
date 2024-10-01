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
	var pattern Pattern
	pattern = [5]LetterColor{Green, Gray, Green, Gray, Yellow}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test2() {
	targetWord := "stare"
	guess := "steer"
	var pattern Pattern
	pattern = [5]LetterColor{Green, Green, Yellow, Gray, Yellow}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test3() {
	targetWord := "sheen"
	guess := "siren"
	var pattern Pattern
	pattern = [5]LetterColor{Green, Gray, Gray, Green, Green}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test4() {
	targetWord := "sheen"
	guess := "elate"
	var pattern Pattern
	pattern = [5]LetterColor{Yellow, Gray, Gray, Gray, Yellow}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test5() {
	targetWord := "messy"
	guess := "sheen"
	var pattern Pattern
	pattern = [5]LetterColor{Yellow, Gray, Yellow, Gray, Gray}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test6() {
	targetWord := "stare"
	guess := "bleed"
	var pattern Pattern
	pattern = [5]LetterColor{Gray, Gray, Yellow, Yellow, Gray}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeFalse)
}

func (this *SolverFixture) Test7() {
	targetWord := "elate"
	guess := "bleed"
	var pattern Pattern
	pattern = [5]LetterColor{Gray, Green, Yellow, Yellow, Gray}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test8() {
	targetWord := "error"
	guess := "revel"
	var pattern Pattern
	pattern = [5]LetterColor{Yellow, Yellow, Gray, Gray, Gray}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test9() {
	targetWord := "freer"
	guess := "error"
	var pattern Pattern
	pattern = [5]LetterColor{Yellow, Green, Gray, Gray, Green}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test10() {
	targetWord := "cacao"
	guess := "cloot"
	var pattern Pattern
	pattern = [5]LetterColor{Green, Gray, Yellow, Gray, Gray}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeTrue)
}
func (this *SolverFixture) Test11() {
	targetWord := "cater"
	guess := "blech"
	var pattern Pattern
	pattern = [5]LetterColor{Gray, Gray, Yellow, Gray, Gray}
	valid := this.Solver.isValidTarget(targetWord, guess, pattern)
	this.So(valid, should.BeFalse)
}

func (this *SolverFixture) TestCalculateExpectedInformationSmart() {
	word := "smart"
	expectedInformation := this.Solver.calculateExpectedInfo(word)
	//this.So(expectedInformation, should.Equal, 5)
	this.So(expectedInformation, should.BeBetween, 5.2, 5.3)
}

func (this *SolverFixture) TestCalculateExpectedInformationRaise() {
	word := "raise"
	expectedInformation := this.Solver.calculateExpectedInfo(word)
	//this.So(expectedInformation, should.Equal, 5)
	this.So(expectedInformation, should.BeBetween, 5.8, 5.9)
}
