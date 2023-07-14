package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("alpha")

	wg.Wait()

	if msg != "alpha" {
		t.Error("Not same")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "gamma"

	// printした段階でos.Stdoutに結果が書き込まれる
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	// たぶんパイプを使ってprintの結果を書き込んで、パイプから結果を読み込んで比較してる
	if !strings.Contains(output, "gamma") {
		t.Errorf("Not equal")
	}

}
