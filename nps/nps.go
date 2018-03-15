package nps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var operators = map[string]bool{"+": true, "-": true, "*": true, "/": true}

func isLegal(a, b, c, d, e, f int) bool {
	var legalMap = map[int]int{100: 1,
		75: 1,
		50: 1,
		25: 1,
		10: 2,
		9:  2,
		8:  2,
		7:  2,
		6:  2,
		5:  2,
		4:  2,
		3:  2,
		2:  2,
		1:  2}

	haveMap := map[int]int{}
	haveMap[a]++
	haveMap[b]++
	haveMap[c]++
	haveMap[d]++
	haveMap[e]++
	haveMap[f]++

	for k, v := range haveMap {
		if v > legalMap[k] {
			return false
		}
	}

	return true
}

func getMoreSolutions(s string) []string {
	newList := []string{s}
	l := strings.Split(s, "")
	if len(l) > 9 {
		return newList
	}
	for k, v := range l {
		if v == "N" {
			var ns []string
			if k != 0 {
				ns = append(ns, l[:k]...)
			}
			ns = append(ns, "NNO")
			if k < len(l) {
				ns = append(ns, l[k+1:]...)
			}
			newList = append(newList, getMoreSolutions(strings.Join(ns, ""))...)
		}
	}

	return newList

}

func legalSolutionList() (s *map[string]bool) {
	baseSolution := "NNO"
	allSolutions := getMoreSolutions(baseSolution)

	solutionMap := make(map[string]bool)

	for _, solution := range allSolutions {

		solutionMap[solution] = true
	}

	return &solutionMap

}

func itrOverOperators(s string, numOperators int) {
	operList := make([]int, numOperators)

	/*

		for i := 0; i < numOperators; i++ {
			sum := s
			for j := 0; j < 4; j++ {
				operList[i] = j

				fmt.Println(operList)

				for _, oper := range operList {
					//fmt.Println(oper)
					operator := ""
					switch oper {
					case 0:
						operator = "+"
					case 1:
						operator = "-"
					case 2:
						operator = "*"
					case 3:
						operator = "/"

					}
					sum = strings.Replace(sum, "O", operator, 1)
				}

			}

			operList[i] = 0

			//fmt.Println(sum)
			//reversePolishSolve(sum)


		}
	*/

}

func itrOverLegalNumbers(numList []int, sol string) {
	numCount := (len(sol) + 1) / 2
	for i := int64(1); i < 64; i++ {
		rpn := strings.Split(sol, "")
		solStr := strings.Join(rpn, " ")
		binRep := fmt.Sprintf("%06b", i)
		letterCount := strings.Count(binRep, "1")

		if letterCount != numCount {
			continue
		}

		for k, v := range binRep {
			number := strconv.Itoa(numList[k])
			if string(v) == "1" {
				solStr = strings.Replace(solStr, "N", number, 1)
			}
		}

		itrOverOperators(solStr, numCount-1)

	}
}

func findAllSolutions(numList []int, solutions *map[string]bool) {
	for sol := range *solutions {
		itrOverLegalNumbers(numList, sol)
	}
}

//Solver will solve a countdown style numbers puzzle.
func Solver(target, a, b, c, d, e, f int) (results []string, err error) {
	if !isLegal(a, b, c, d, e, f) {
		return results, errors.New("Number combination is not legal for Countdown")
	}

	solutions := legalSolutionList()

	numList := []int{a, b, c, d, e, f}

	findAllSolutions(numList, solutions)

	return results, err
}

func reversePolishToBrackets(rpn []string) (s []string) {
	if len(rpn) > 1 {
		for k, v := range rpn {
			if _, ok := operators[v]; ok {
				s := "(" + rpn[k-2] + v + rpn[k-1] + ")"
				var ns []string
				ns = append(ns, rpn[:k-2]...)
				ns = append(ns, s)
				if k < len(rpn)-1 {
					ns = append(ns, rpn[k+1:]...)
				}
				return reversePolishToBrackets(ns)
			}
		}
	}

	return rpn
}
