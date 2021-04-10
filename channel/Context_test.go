package channel

import (
	"context"
	"fmt"
)

func PlusOneCtx(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func Example_PlusOneCtx() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	nums := PlusOneCtx(ctx, PlusOneCtx(ctx, PlusOneCtx(ctx, PlusOneCtx(ctx, PlusOneCtx(ctx, c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			cancel()
			break
		}
	}
	// Output:
	// .
}
