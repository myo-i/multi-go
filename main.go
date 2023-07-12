package main

import (
	"fmt"
	"sync"
)

func print(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup

	count := []string{
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
	}

	wg.Add(len(count))
	for i, x := range count {
		go print(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)
	// Secondのみ出力
	print("Second", &wg)
}
