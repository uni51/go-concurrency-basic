package main

/*
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("1st Goroutine Start")
		time.Sleep(1 * time.Second)
		fmt.Println("1st Goroutine Done")
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("2nd Goroutine Start")
		time.Sleep(1 * time.Second)
		fmt.Println("2nd Goroutine Done")
	}()

	wg.Wait()
}
*/
