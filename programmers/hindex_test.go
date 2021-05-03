package programmers

import "fmt"

func Example_solution_hindex() {
	testCitations := []int{3, 0, 6, 1, 5}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{0, 1, 2}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{1, 4}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{12, 11, 10, 9, 8, 1}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{6, 6, 6, 6, 6, 1}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{20, 21, 22, 23}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{4, 4, 4}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{5, 5, 5, 5}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{0, 1, 1, 1, 1, 3, 3, 4}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{2, 2, 2, 2, 2}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{22, 42}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{2}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{0, 0, 0}
	fmt.Println(solution_hindex(testCitations))
	testCitations = []int{0, 1, 1}
	fmt.Println(solution_hindex(testCitations))
	// Output:
	// .
}

func solution_hindex(citations []int) int {
	maxIndex := 0
	hIndex := 0
	tmpIndex := map[int]int{}
	for _, c := range citations {
		if maxIndex < c {
			maxIndex = c
		}
	}
	if maxIndex < len(citations) {
		maxIndex = len(citations)
	}
	for i := 1; i < maxIndex; i++ {
		cnt := 0
		for _, c := range citations {
			if c >= i {
				cnt++
			}
		}
		if i <= cnt {
			hIndex = i
		}
		tmpIndex[i] = cnt
	}
	if len(tmpIndex) == 0 {
		return 0
	}
	return hIndex
}
