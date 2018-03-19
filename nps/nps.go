package nps

import (
	"errors"
	"fmt"
	"math"
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

func solve(s string) float64 {
	equation := strings.Split(s, " ")
	if len(equation) == 1 {
		a, _ := strconv.ParseFloat(equation[0], 64)
		return float64(a)
	}
	var res float64
	for k, v := range equation {
		var a float64
		var b float64
		if v == "+" {
			a, _ = strconv.ParseFloat(equation[k-2], 64)
			b, _ = strconv.ParseFloat(equation[k-1], 64)
			res = float64(a) + float64(b)
		} else if v == "-" {
			a, _ = strconv.ParseFloat(equation[k-2], 64)
			b, _ = strconv.ParseFloat(equation[k-1], 64)
			res = float64(a) - float64(b)

		} else if v == "*" {
			a, _ = strconv.ParseFloat(equation[k-2], 64)
			b, _ = strconv.ParseFloat(equation[k-1], 64)
			res = float64(a) * float64(b)
		} else if v == "/" {
			a, _ = strconv.ParseFloat(equation[k-2], 64)
			b, _ = strconv.ParseFloat(equation[k-1], 64)
			res = float64(a) / float64(b)
		} else {
			continue
		}

		var ne []string
		ne = append(ne, equation[:k-2]...)
		ne = append(ne, strconv.FormatFloat(res, 'E', -1, 64))
		ne = append(ne, equation[k+1:]...)

		return solve(strings.Join(ne, " "))
	}

	return 0
}

func itrOverOperators(s string, numOperators int) (m map[string][]string) {
	m = make(map[string][]string)
	maxBytes := math.Pow(2, float64(numOperators*2))

	var br string
	switch numOperators {
	case 1:
		br = "%02b"
	case 2:
		br = "%04b"
	case 3:
		br = "%06b"
	case 4:
		br = "%08b"
	case 5:
		br = "%010b"
	}

	for i := 0; i < int(maxBytes); i++ {
		rpn := s
		b := fmt.Sprintf(br, i)
		opers := make([]string, numOperators)
		for j := 0; j < len(b); j += 2 {
			switch oper := b[j : j+2]; oper {
			case "00":
				opers[(j+1)/2] = "+"
			case "01":
				opers[(j+1)/2] = "-"
			case "10":
				opers[(j+1)/2] = "*"
			case "11":
				opers[(j+1)/2] = "/"
			}

		}

		for k := 0; k < numOperators; k++ {
			rpn = strings.Replace(rpn, "O", opers[k], 1)
		}

		res := solve(rpn)
		resstr := strconv.FormatFloat(res, 'f', -1, 64)
		brackets := strings.Join(reversePolishToBrackets(strings.Split(rpn, " ")), " ")
		fmt.Println(resstr)
		fmt.Println(brackets)
		m[resstr] = append(m[resstr], brackets)
	}

	return m
}

func itrOverLegalNumbers(numList []int, sol string) (m map[string][]string) {
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

		return itrOverOperators(solStr, numCount-1)

	}

	return m

}

func findAllSolutions(numList []int, solutions *map[string]bool) (m map[string][]string) {
	for sol := range *solutions {
		return itrOverLegalNumbers(numList, sol)
	}

	return m
}

//Solver will solve a countdown style numbers puzzle.
func Solver(a, b, c, d, e, f int) (m map[string][]string, err error) {
	if !isLegal(a, b, c, d, e, f) {
		return m, errors.New("Number combination is not legal for Countdown")
	}

	solutions := legalSolutionList()

	numList := []int{a, b, c, d, e, f}

	return findAllSolutions(numList, solutions), nil

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
