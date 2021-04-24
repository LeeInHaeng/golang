package programmers

import "fmt"

func Example_solution_nejuk() {
	testA := []int{1, 2, 3, 4}
	testB := []int{-3, -1, 0, 2}
	fmt.Println(solution_nejuk(testA, testB))
	// Output:
	// .
}

func solution_nejuk(a []int, b []int) int {
	answer := 0

	for idx, _ := range a {
		answer += a[idx] * b[idx]
	}

	return answer
}
