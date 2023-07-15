package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {

	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Done()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Done()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Done()

	printMessage()
}
