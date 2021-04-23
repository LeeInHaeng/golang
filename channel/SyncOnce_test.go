package channel

import (
	"fmt"
	"sync"
)

func Example_SyncOnce() {
	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("Initialized")
			})
			fmt.Println("Go", i)
		}(i)
	}
	wg.Wait()
	// Output:
	// .
}
