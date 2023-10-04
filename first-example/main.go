package main

import (
	"fmt"
	"sync"
)

func printSomething(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(str)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gama",
		"delta",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	// when compiler sees anything after 'go' commnad executed in its own goRoutine. SO its creates a new goRoutine and hand it to the go-scheduler
	for i, x := range words {
		go printSomething(fmt.Sprintf("%d : %s", i, x), &wg)
	}

	wg.Wait()

	// time.Sleep(1 * time.Second)
	wg.Add(1)

	printSomething("This is the second thing to be printed", &wg)

}
