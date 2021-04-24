package programmers

import (
	"fmt"
	"strings"
)

func Example_solution_midletter() {
	testCase := "abcde"
	fmt.Println(solution_midletter(testCase))
	testCase = "qwer"
	fmt.Println(solution_midletter(testCase))
	// Output:
	// .
}

func solution_midletter(s string) string {
	posFloat := float64(len(s)) / 2
	posInt := len(s) / 2
	splitString := strings.Split(s, "")

	if posFloat != float64(posInt) {
		return splitString[posInt]
	} else {
		return splitString[posInt-1] + splitString[posInt]
	}
}
