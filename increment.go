package main

import (
	"fmt"
	"sync"
)

var value int
var mu sync.Mutex
var wg = sync.WaitGroup{}

func increment() {
	for i := 0; i < 100; i++ {
		mu.Lock()
		value++
		mu.Unlock()
	}
	wg.Done()
}

func main() {

	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go increment()

	}
	wg.Wait()
	fmt.Println(value)
}
