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

func (this *SolverFixture) Test() {
	targetWord := "snake"
	turn := Turn{
		Guess:   "slain",
		Pattern: Pattern{Green, Gray, Green, Gray, Yellow},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test2() {
	targetWord := "stare"
	turn := Turn{
		Guess:   "steer",
		Pattern: Pattern{Green, Green, Yellow, Gray, Yellow},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test3() {
	targetWord := "sheen"
	turn := Turn{
		Guess:   "siren",
		Pattern: Pattern{Green, Gray, Gray, Green, Green},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test4() {
	targetWord := "sheen"
	turn := Turn{
		Guess:   "elate",
		Pattern: Pattern{Yellow, Gray, Gray, Gray, Yellow},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test5() {
	targetWord := "messy"
	turn := Turn{
		Guess:   "sheen",
		Pattern: Pattern{Yellow, Gray, Yellow, Gray, Gray},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test6() {
	targetWord := "stare"
	turn := Turn{
		Guess:   "bleed",
		Pattern: Pattern{Gray, Gray, Yellow, Yellow, Gray},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeFalse)
}

func (this *SolverFixture) Test7() {
	targetWord := "elate"
	turn := Turn{
		Guess:   "bleed",
		Pattern: Pattern{Gray, Green, Yellow, Yellow, Gray},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test8() {
	targetWord := "error"
	turn := Turn{
		Guess:   "revel",
		Pattern: Pattern{Yellow, Yellow, Gray, Gray, Gray},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test9() {
	targetWord := "freer"
	turn := Turn{
		Guess:   "error",
		Pattern: Pattern{Yellow, Green, Gray, Gray, Green},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeTrue)
}

func (this *SolverFixture) Test10() {
	targetWord := "cacao"
	turn := Turn{
		Guess:   "cloot",
		Pattern: Pattern{Green, Gray, Yellow, Gray, Gray},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeTrue)
}
func (this *SolverFixture) Test11() {
	targetWord := "cater"
	turn := Turn{
		Guess:   "blech",
		Pattern: Pattern{Gray, Gray, Yellow, Gray, Gray},
	}
	valid := this.Solver.isValidTarget(targetWord, turn)
	this.So(valid, should.BeFalse)
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

func (this *SolverFixture) TestMaximizeExpectedInformationFirstGuess() {
	firstGuess := this.Solver.maximizeExpectedInformation()
	this.So(firstGuess, should.Equal, "soare")
}

func (this *SolverFixture) TestSingleGame() {
	numGuesses, err := this.Evaluator.PlayGame("angry", this.Solver)
	this.So(err, should.BeNil)
	this.So(numGuesses, should.BeLessThanOrEqualTo, MaxNumGuesses)
}
