package main

import (
	"fmt"
	"runtime"
	"time"
)

func DoSomething(done chan bool) {
	// Do Stuff Here
	time.Sleep(time.Second)

	select {
	case done <- true:
	default:
		return
	}
}

func Timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func main() {
	for i := 0; i < 10; i++ {
		err := Timeout(DoSomething)
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}
