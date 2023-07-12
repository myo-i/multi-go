package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_print(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	// printした段階でos.Stdoutに結果が書き込まれる
	go print("One", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	// たぶんパイプを使ってprintの結果を書き込んで、パイプから結果を読み込んで比較してる
	if !strings.Contains(output, "One") {
		t.Errorf("Not equal")
	}
}
