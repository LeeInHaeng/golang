package programmers

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type SameBit struct {
	Idx int
	Num int64
}

func Example_solution_twonotsamebit() {
	// 10 , 11 시간초과
	testNumbers := []int64{2, 7}
	fmt.Println(solution_twonotsamebit(testNumbers))
	// Output:
	// .
}

func solution_twonotsamebit(numbers []int64) []int64 {
	answer := make([]int64, len(numbers))
	for resBit := range getMinTwoNotSameBitGo(numbers) {
		answer[resBit.Idx] = resBit.Num
	}
	return answer
}

func getMinTwoNotSameBitGo(numbers []int64) <-chan SameBit {
	chanRes := make(chan SameBit)
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for idx, number := range numbers {
		req := SameBit{idx, number}
		go func(req SameBit) {
			defer wg.Done()
			chanRes <- getMinTwoNotSameBit(req)
		}(req)
	}

	go func() {
		wg.Wait()
		close(chanRes)
	}()

	return chanRes
}

func getMinTwoNotSameBit(req SameBit) SameBit {
	startBinary := strconv.FormatInt(req.Num, 2)
	startBit := strings.Split(startBinary, "")
	startBitLen := len(startBit)

	for true {
		req.Num += 1
		compareBinary := strconv.FormatInt(req.Num, 2)
		compareBit := strings.Split(compareBinary, "")
		compareBitLen := len(compareBit)

		compareCnt := 0
		if startBitLen != compareBitLen {
			compareCnt += compareBitLen - startBitLen
		}

		if compareCnt <= 2 {
			startIdx := 0
			for i := compareCnt; i < compareBitLen; i++ {
				if startBit[startIdx] != compareBit[i] {
					compareCnt++
				}
				startIdx++
				if compareCnt > 2 {
					break
				}
			}

			if compareCnt <= 2 {
				break
			}
		}
	}
	return SameBit{req.Idx, req.Num}
}
