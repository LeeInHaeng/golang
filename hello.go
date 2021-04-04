package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello World!")

	myInts := []int{50, 30, 10, 20}
	sort.Ints(myInts)
	fmt.Println("sorted : ", myInts)
}
