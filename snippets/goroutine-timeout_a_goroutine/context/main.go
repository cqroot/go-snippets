package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func DoSomething(ctx context.Context) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Cancel goroutine by context")
			return
		default:
			// Do Stuff Here
			time.Sleep(time.Microsecond * 100)
		}
	}
}

func main() {
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		go DoSomething(ctx)
	}

	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
}
