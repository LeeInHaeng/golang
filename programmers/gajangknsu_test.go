package programmers

import (
	"sort"
	"strconv"
)

func Example_gajangknsu_solution(numbers []int) string {
	sort.Slice(numbers, func(i, j int) bool {
		s1 := strconv.Itoa(numbers[i])
		s2 := strconv.Itoa(numbers[j])
		return s1+s2 > s2+s1
	})

	if numbers[0] == 0 {
		return "0"
	}
	answer := ""
	for i := range numbers {
		answer += strconv.Itoa(numbers[i])
	}
	return answer
}
