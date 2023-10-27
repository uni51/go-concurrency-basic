package main

// import "sync"

/*
func main() {
	var wg sync.WaitGroup

	say := "Hello"

	wg.Add(1)

	go func() {
		defer wg.Done()
		say = "Good Bye"
	}()

	wg.Wait()

	fmt.Println(say) // Good Bye
}
*/

/*
func main() {
	var wg sync.WaitGroup

	tasks := []string{"A", "B", "C"}

	for _, task := range tasks {
		wg.Add(1)

		go func(task string) {
			defer wg.Done()
			println(task) // C B A
		}(task)
	}

	wg.Wait()
}
*/
