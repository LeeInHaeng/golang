package channel

import (
	"fmt"
	"sync"
)

func Min(nums []int) int {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}

/*
func ParallelMin(nums []int, divide int) int {
	idx := 0
	chanMins := make([]int, divide)
	for chanMin := range ParallelMinGo(nums, divide) {
		chanMins[idx] = chanMin
		idx++
	}
	return Min(chanMins)
}

func ParallelMinGo(nums []int, divide int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < divide; i++ {
			start := i * divide
			end := (i + 1) * divide
			if i == divide-1 {
				end = len(nums)
			}
			c <- Min(nums[start:end])
		}
	}()
	return c
}
*/

func ParallelMinGo(nums []int, divide int) <-chan int {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(divide)
	for i := 0; i < divide; i++ {
		start := i * divide
		end := (i + 1) * divide
		if i == divide-1 {
			end = len(nums)
		}
		go func(int, int) {
			defer wg.Done()
			c <- Min(nums[start:end])
		}(start, end)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}

func ParallelMin(nums []int, divide int) int {
	mins := make([]int, divide)
	idx := 0
	for chanMins := range ParallelMinGo(nums, divide) {
		mins[idx] = chanMins
		idx++
	}
	return Min(mins)
}

func ExampleMin() {
	fmt.Println(Min([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
	}))
	fmt.Println(ParallelMin([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
	}, 4))
	// Output:
	// 23
	// 23
}
