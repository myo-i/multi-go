package producerconsumer

import (
	"image/color"
	"math/rand"
	"time"
)

const NumberOfChicken = 10

var chickenMade, chickenFailed, total int

type Producer struct {
	data chan ChickenOrder
	quit chan chan error
}

type ChickenOrder struct {
	chickenNumber int
	message       string
	success       bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func kfcShop(chickenMaker *Producer) {
	// チャネルに何かしらの情報を取得するまでは起動し続ける
	for {
		// 処理
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Cyan("Open KFC!!")
	color.Cyan("----------")

	// create producer
	kfcjob := &Producer{
		data: make(chan ChickenOrder),
		quit: make(chan chan error),
	}

	// run producer in the background
	go kfcShop()
}
