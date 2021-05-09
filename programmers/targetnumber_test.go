package programmers

import "fmt"

func Example_solution_targetnumber() {
	testNumbers := []int{1, 1, 1, 1, 1}
	testTarget := 3
	fmt.Println(solution_targetnumber(testNumbers, testTarget))
	// Output:
	// .
}

func solution_targetnumber(numbers []int, target int) int {
	cnt := 0
	for res := range recurTargetNumberGo(numbers) {
		if res == target {
			cnt++
		}
	}

	return cnt
}

func recurTargetNumberGo(numbers []int) <-chan int {
	resChan := make(chan int)

	go func(resChan chan int) {
		defer close(resChan)
		recurTargetNumber(0, numbers, 0, resChan)
	}(resChan)

	return resChan
}

func recurTargetNumber(calcNumber int, numbers []int, idx int, resChan chan int) {
	if len(numbers)-1 == idx {
		posResNumber := calcNumber + numbers[idx]
		resChan <- posResNumber
		negaResNumber := calcNumber - numbers[idx]
		resChan <- negaResNumber
		return
	}

	posCalcNumber := calcNumber + numbers[idx]
	recurTargetNumber(posCalcNumber, numbers, idx+1, resChan)
	negaCalcNumber := calcNumber - numbers[idx]
	recurTargetNumber(negaCalcNumber, numbers, idx+1, resChan)
}
