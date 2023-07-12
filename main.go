package main

import "fmt"

func print(s string) {
	fmt.Println(s)
}

func main() {
	// Secondのみ出力
	go print("First")
	print("Second")
}
