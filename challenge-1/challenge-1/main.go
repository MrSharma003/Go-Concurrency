package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello from print")
	fmt.Println(msg)
}

func another() {
	a := 3*2 + 0 - 1
	// fmt.Println("Hello from another")
	a = a + 1
	fmt.Println("Hello from another")
	fmt.Println("Hello from another again")
}

func main() {
	var wg sync.WaitGroup

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	wg.Add(2)

	go updateMessage("Hello, universe!", &wg)
	go printMessage(&wg)
	another()
	wg.Wait()

	// wg.Add(1)
	// go updateMessage("Hello, cosmos!", &wg)
	// wg.Wait()
	// printMessage()
	// // wg.Wait()

	// wg.Add(1)
	// go updateMessage("Hello, world!", &wg)
	// wg.Wait()
	// printMessage()

	// wg.Wait()
}
