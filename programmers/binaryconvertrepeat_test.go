package programmers

import (
	"fmt"
	"strconv"
	"strings"
)

func Example_solution_binaryconvertrepeat() {
	testS := "110010101001"
	fmt.Println(solution_binaryconvertrepeat(testS))
	testS = "01110"
	fmt.Println(solution_binaryconvertrepeat(testS))
	testS = "1111111"
	fmt.Println(solution_binaryconvertrepeat(testS))
	// Output:
	// .
}

func solution_binaryconvertrepeat(s string) []int {
	sliceS := strings.Split(s, "")
	answerRemovedZero := 0
	answerCnt := 0

	for true {
		answerCnt++
		removedZero := 0

		stringLen := len(sliceS)
		if stringLen == 1 {
			break
		}
		for _, s := range sliceS {
			if s == "0" {
				removedZero++
			}
		}
		answerRemovedZero += removedZero
		binaryS := strconv.FormatInt(int64(stringLen-removedZero), 2)
		sliceS = strings.Split(binaryS, "")
	}
	return []int{answerCnt - 1, answerRemovedZero}
}
