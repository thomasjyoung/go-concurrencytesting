package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Program Start")
	// launch 2 additional go routines
	fmt.Println("Our Go Routines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("I am go routine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("I am go routine 2")
		wg.Done()
	}()
	fmt.Println("Our Go Routines end", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Program End")
}
