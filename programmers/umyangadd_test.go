package programmers

import "fmt"

func Example_solution_umyangadd() {
	testAbsolutes := []int{4, 7, 12}
	testSigns := []bool{true, false, true}
	fmt.Println(solution_umyangadd(testAbsolutes, testSigns))
	testAbsolutes = []int{1, 2, 3}
	testSigns = []bool{false, false, true}
	fmt.Println(solution_umyangadd(testAbsolutes, testSigns))
	// Output:
	// .
}

func solution_umyangadd(absolutes []int, signs []bool) int {
	answer := 0

	for idx, sign := range signs {
		if sign == true {
			answer += absolutes[idx]
		} else {
			answer -= absolutes[idx]
		}
	}
	return answer
}
