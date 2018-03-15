package nps

import (
	"fmt"
	"strings"
)

var operators = map[string]bool{"+": true, "-": true, "*": true, "/": true}

//Solver will solve a countdown style numbers puzzle.
//func Solver(a, b, c, d, e, f int) (results []string, err error) {
//
//}

func reversePolishToBrackets(rpn string) (s string) {
	fmt.Println(rpn)
	ignore := 0
	lastNumPos := 0
	startIgnore := 0
	endIgnore := 0
	var intermediateString []string
	for k, v := range rpn {
		if string(v) == "(" {
			ignore++
			if ignore == 1 {
				startIgnore = k
			}
		} else if string(v) == ")" {
			ignore--
			if ignore == 0 {
				endIgnore = k
			}
		}

		if ignore == 0 {
			lastNumPos = k

			if _, ok := operators[string(v)]; ok {
				ns := "(" + rpn[lastNumPos] + string(v) + rpn[startIgnore:endIgnore] + ")"
				if k < len(sl)-1 {
					ns = append(ns, sl[k:]...)
				}

				return reversePolishToBrackets(strings.Join(ns, " "))
			}
		} else {
			intermediateString = append(intermediateString, v)
		}

	}
	return s
}
