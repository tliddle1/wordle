package main

import (
	"fmt"

	"github.com/tliddle1/wordle"
	"github.com/tliddle1/wordle/pkg/solver"
)

func main() {
	evaluator := wordle.NewEvaluator()
	thomasSolver := solver.NewThomasSolver()
	avgNumGuesses, err := evaluator.EvaluateSolver(thomasSolver)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(avgNumGuesses)
}
