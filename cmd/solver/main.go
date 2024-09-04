package main

import (
	"fmt"

	"github.com/tliddle1/wordle"
	"github.com/tliddle1/wordle/pkg/solver"
)

func main() {
	evaluator := wordle.NewEvaluator()
	solver := solver.NewThomasSolver()
	avgNumGuesses, err := evaluator.EvaluateSolver(solver)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(avgNumGuesses)
}
