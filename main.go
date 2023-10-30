package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello from %v!\n", id)
}

func main() {
	var wg sync.WaitGroup

	var CPU int = runtime.NumCPU()

	// wg.Add(CPU)

	for i := 0; i < CPU; i++ {
		wg.Add(1)
		go Hello(&wg, i)
	}

	wg.Wait()
}
