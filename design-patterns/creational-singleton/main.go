package main

import (
	"fmt"
	"sync"

	"github.com/cqroot/go-snippets/design-patterns/creational-singleton/singleton"
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
