package channel

import "fmt"

func Example_recurChan() {
	for res := range simpleRecurGo() {
		fmt.Println(res)
	}
	// Output:
	// .
}

func simpleRecurGo() <-chan int {
	resChan := make(chan int)

	go func(resChan chan int) {
		defer close(resChan)
		simpleRecur(5, resChan)
	}(resChan)

	return resChan
}

func simpleRecur(num int, resChan chan int) {
	if num < 0 {
		resChan <- num
		return
	}
	num -= 1
	simpleRecur(num, resChan)
}
