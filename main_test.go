package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello")
	go updateMessage("12345")
	wg.Wait()

	if msg != "Hello" {
		t.Errorf("Failed to test")
	}

}
