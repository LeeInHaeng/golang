package programmers

import "fmt"

func Example_solution_cheukbok() {
	testCaseN := 5
	testCaseLost := []int{2, 4}
	testCaaseReserve := []int{1, 3, 5}
	fmt.Println(solution_cheukbok(testCaseN, testCaseLost, testCaaseReserve))
	testCaseN = 4
	testCaseLost = []int{3, 1}
	testCaaseReserve = []int{2, 4}
	fmt.Println(solution_cheukbok(testCaseN, testCaseLost, testCaaseReserve))
	// Output:
	// .
}

func solution_cheukbok(n int, lost []int, reserve []int) int {
	students := make([]int, n)
	for idx, _ := range students {
		students[idx] = 1
	}
	for _, l := range lost {
		students[l-1] -= 1
	}
	for _, r := range reserve {
		students[r-1] += 1
	}

	for i := 0; i < n; i++ {
		if students[i] == 0 {
			if i != 0 && students[i-1] > 1 {
				students[i-1] -= 1
				students[i] += 1
			} else if i+1 != n && students[i+1] > 1 {
				students[i+1] -= 1
				students[i] += 1
			}
		}
	}

	cnt := 0
	for _, s := range students {
		if s > 0 {
			cnt++
		}
	}

	return cnt
}
