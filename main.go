package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {

	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, cat")
	go updateMessage("Hello, bird")
	wg.Wait()

	fmt.Println(msg)

}
