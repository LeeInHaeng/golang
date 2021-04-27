package channel

import (
	"fmt"
	"sync"
	"time"
)

func sleepGo() <-chan string {
	myChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		myChan <- "sleep 3"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 6)
		myChan <- "sleep 6"
	}()

	go func() {
		wg.Wait()
		close(myChan)
	}()

	return myChan
}

func normalSleep() {
	time.Sleep(time.Second * 3)
	time.Sleep(time.Second * 6)
}

func Example_sleepGo() {
	/*
		select {
		case res := <-sleepGo():
			fmt.Println(res) // sleep 3 (수행시간 3)
		}
	*/

	for res2 := range sleepGo() {
		fmt.Println(res2) // sleep 3 sleep 6 (수행시간 6)
	}
	// Output:
	// .
}
