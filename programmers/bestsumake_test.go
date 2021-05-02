package programmers

import (
	"fmt"
)

func Example_solution_bestsumake() {
	testNumber := "1924"
	testK := 2
	fmt.Println(solution_bestsumake(testNumber, testK))
	testNumber = "1231234"
	testK = 3
	fmt.Println(solution_bestsumake(testNumber, testK))
	testNumber = "656516565"
	testK = 1
	fmt.Println(solution_bestsumake(testNumber, testK))
	testNumber = "1203045"
	testK = 2
	fmt.Println(solution_bestsumake(testNumber, testK))
	// Output:
	// .
}

func solution_bestsumake(number string, k int) string {
	// O(n^2) 로 시간초과
	// 재귀연습
	maxNum := "0"
	for nums := range getBestNumGo(number, k) {
		if maxNum < nums {
			maxNum = nums
		}
	}
	return maxNum
}

func getBestNumGo(number string, k int) <-chan string {
	chanRes := make(chan string)
	numLen := len(number) - k

	go func(chanRes chan string) {
		defer close(chanRes)
		getBestNumRecur(chanRes, "", number, numLen)
	}(chanRes)

	return chanRes
}

func getBestNumRecur(resChan chan string, primeNum string, extraNum string, numLen int) {
	if len(primeNum) > numLen {
		return
	}
	for idx, snum := range extraNum {
		mergeNum := primeNum + string(snum)
		if len(mergeNum) == numLen {
			resChan <- mergeNum
		}
		getBestNumRecur(resChan, mergeNum, extraNum[idx+1:], numLen)
	}
}
