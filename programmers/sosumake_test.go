package programmers

import (
	"fmt"
	"sort"
	"sync"
)

func Example_solution_sosumake() {
	testCase := []int{1, 2, 3, 4}
	fmt.Println(solution_sosumake(testCase))
	testCase = []int{1, 2, 7, 6, 4}
	fmt.Println(solution_sosumake(testCase))
	// Output:
	// .
}

func solution_sosumake(nums []int) int {

	answer := 0
	for goAnswer := range getSosuMakeGo(nums) {
		answer = goAnswer
	}

	return answer
}

func getSosuMakeGo(nums []int) <-chan int {
	var allNums []int
	var allPrimes []int
	var wg sync.WaitGroup
	answer := 0
	chanAnswer := make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		allNums = getAllNums(nums)
	}()
	go func() {
		defer wg.Done()
		allPrimes = getAllPrimes(nums)
	}()

	go func() {
		wg.Wait()
		for _, num := range allNums {
			if allPrimes[num] == 0 {
				answer++
			}
		}
		chanAnswer <- answer
		close(chanAnswer)
	}()
	return chanAnswer
}

func getAllNums(nums []int) []int {
	allNums := make([]int, 1)

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				allNums = append(allNums, nums[i]+nums[j]+nums[k])
			}
		}
	}
	return allNums
}

func getAllPrimes(nums []int) []int {
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	sort.Ints(copyNums)
	maxNum := copyNums[len(nums)-1] + copyNums[len(nums)-2] + copyNums[len(nums)-3]
	allPrimes := make([]int, maxNum+1)
	allPrimes[0] = 1
	allPrimes[1] = 1

	for i := 2; i < len(allPrimes); i++ {
		if allPrimes[i] == 0 {
			for j := 2; j*i < len(allPrimes); j++ {
				allPrimes[j*i] = 1
			}
		}
	}

	return allPrimes
}
