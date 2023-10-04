package main

import (
	"fmt"
	"strings"
)

func pingPong(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go pingPong(ping, pong)

	for {
		fmt.Print("->")
		var user string
		fmt.Scanln(&user)

		if "q" == strings.ToLower(user) {
			break
		}

		ping <- user
		response := <-pong
		fmt.Println("Response", response)
	}
	fmt.Println("All done. Closing Channels")
	close(ping)
	close(pong)
}
