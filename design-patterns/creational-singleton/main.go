package main

import (
	"fmt"
	"go-snippets/design-patterns/creational-singleton/singleton"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println(singleton.New().Time)
		}()
	}

	wg.Wait()
}
