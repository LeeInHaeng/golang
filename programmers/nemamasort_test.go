package programmers

import (
	"fmt"
	"sort"
	"strings"
)

type MapString struct {
	SortKey string
	String  string
}

func Example_solution_nemamsort() {
	testStrings := []string{"sun", "bed", "car"}
	testN := 1
	fmt.Println(solution_nemamsort(testStrings, testN))
	testStrings = []string{"abce", "abcd", "cdx"}
	testN = 2
	fmt.Println(solution_nemamsort(testStrings, testN))
	testStrings = []string{"abzcd", "cdzab", "abzfg", "abzaa", "abzbb", "bbzaa"}
	testN = 2
	fmt.Println(solution_nemamsort(testStrings, testN))
	// Output:
	// .
}

func solution_nemamsort(myStrings []string, n int) []string {
	sortMapString := make([]MapString, len(myStrings))
	for idx, myString := range myStrings {
		splitString := strings.Split(myString, "")
		sortMapString[idx] = MapString{splitString[n], myString}
	}

	sort.Slice(sortMapString, func(i, j int) bool {
		if sortMapString[i].SortKey == sortMapString[j].SortKey {
			return sortMapString[i].String < sortMapString[j].String
		}
		return sortMapString[i].SortKey < sortMapString[j].SortKey
	})

	answer := make([]string, len(sortMapString))
	for idx, sorted := range sortMapString {
		answer[idx] = sorted.String
	}

	return answer
}
