package nps

import (
	"testing"
)

func TestReversePolishToBrackets(t *testing.T) {
	res := reversePolishToBrackets([]string{"6", "75", "*", "50", "/", "100", "3", "+", "*", "25", "+"})

	if res[0] != ("((((6*75)/50)*(100+3))+25)") {
		t.Fail()
	}
}

func TestIsLegal(t *testing.T) {
	if !isLegal(1, 2, 3, 4, 5, 6) {
		t.Fail()
	}

	if isLegal(100, 100, 1, 1, 1, 1) {
		t.Fail()
	}
}

func TestLegalSolutionsList(t *testing.T) {
	solutions := legalSolutionList()
	if len(*solutions) != 64 {
		t.Fail()
	}
}

func TestSolver(t *testing.T) {
	Solver(999, 100, 75, 1, 2, 3, 4)
}
