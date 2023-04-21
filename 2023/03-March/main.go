package main

import (
	"fmt"
	"sync"
)

func main() {
	var i int
	defer fmt.Println(i)

	var wg sync.WaitGroup
	for i < 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			i++
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
