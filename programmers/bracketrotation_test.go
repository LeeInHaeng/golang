package programmers

import (
	"fmt"
	"strings"
	"sync"
)

func Example_solution_bracketrotation() {
	testS := "[](){}"
	fmt.Println(solution_bracketrotation(testS))
	testS = "}]()[{"
	fmt.Println(solution_bracketrotation(testS))
	testS = "[)(]"
	fmt.Println(solution_bracketrotation(testS))
	testS = "}}}"
	fmt.Println(solution_bracketrotation(testS))
	testS = "("
	fmt.Println(solution_bracketrotation(testS))
	// Output:
	// .
}

func solution_bracketrotation(s string) int {
	cnt := 0
	for chanRes := range rotationBracketGo(s) {
		if chanRes == true {
			cnt++
		}
	}
	return cnt
}

func rotationBracketGo(s string) <-chan bool {
	chanRes := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(len(s))

	for i := 0; i < len(s); i++ {
		rotatedString := s[i:] + s[:i]
		go func(reqString string) {
			defer wg.Done()
			chanRes <- isValidBracket(reqString)
		}(rotatedString)
	}

	go func() {
		wg.Wait()
		close(chanRes)
	}()

	return chanRes
}

func isValidBracket(reqString string) bool {
	s := strings.Split(reqString, "")
	var stack []string

	if s[0] == "]" || s[0] == ")" || s[0] == "}" {
		return false
	}
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if s[i] == "[" || s[i] == "(" || s[i] == "{" {
			stack = append(stack, s[i])
		} else {
			popIdx := len(stack) - 1
			if popIdx < 0 && (s[i] == "]" || s[i] == ")" || s[i] == "}") {
				return false
			}
			if s[i] == "]" {
				if stack[popIdx] != "[" {
					return false
				} else {
					stack = stack[:popIdx]
				}
			} else if s[i] == ")" {
				if stack[popIdx] != "(" {
					return false
				} else {
					stack = stack[:popIdx]
				}
			} else if s[i] == "}" {
				if stack[popIdx] != "{" {
					return false
				} else {
					stack = stack[:popIdx]
				}
			}
		}
	}
	if len(stack) != 0 {
		return false
	}

	return true
}
