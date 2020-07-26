package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Have a var to hold incrementer value
	counter := 0

	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			holderValue := counter
			runtime.Gosched()
			holderValue++
			counter = holderValue
			fmt.Println(counter)
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
