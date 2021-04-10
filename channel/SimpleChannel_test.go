package channel

import "fmt"

func Example_simpleChannel() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()
	for num := range c {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
}
