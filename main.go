package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()

// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }
// func main() {
// 	msg = "Hello, world!"

// 	var mutex sync.Mutex

// 	wg.Add(2)
// 	go updateMessage("Hello, cat", &mutex)
// 	go updateMessage("Hello, bird", &mutex)
// 	wg.Wait()

// 	fmt.Println(msg)

// }
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
