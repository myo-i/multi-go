package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
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

func makeChicken(chickenNumber int) *ChickenOrder {
	chickenNumber++
	if chickenNumber <= NumberOfChicken {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order number #%d\n", chickenNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			chickenFailed++
		} else {
			chickenMade++
		}
		total++

		fmt.Printf("Making chicken #%d. It will take %d seconds....\n", chickenNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** out of ingredients for chicken #%d", chickenNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** quit while making chicken #%d", chickenNumber)
		} else {
			success = true
			msg = fmt.Sprintf("#%d is reafy", chickenNumber)
		}

		c := ChickenOrder{
			chickenNumber: chickenNumber,
			message:       msg,
			success:       success,
		}

		return &c
	}

	return &ChickenOrder{
		chickenNumber: chickenNumber,
	}
}

func kfcShop(chickenMaker *Producer) {
	// チャネルに何かしらの情報を取得するまでは起動し続ける
	var i = 0
	for {
		currentChicken := makeChicken(i)
		if currentChicken != nil {
			i = currentChicken.chickenNumber
			select {
			case chickenMaker.data <- *currentChicken:

			case quitChan := <-chickenMaker.quit:
				close(chickenMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	// rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))

	color.Cyan("KFC is open!!")
	color.Cyan("----------")

	// create producer
	kfcjob := &Producer{
		data: make(chan ChickenOrder),
		quit: make(chan chan error),
	}

	// run producer in the background
	go kfcShop(kfcjob)

	for i := range kfcjob.data {
		if i.chickenNumber <= NumberOfChicken {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery", i.chickenNumber)
			} else {
				color.Red(i.message)
				color.Red("Can't delivery!!")
			}
		} else {
			color.Cyan("Done making chickens...")
			err := kfcjob.Close()
			if err != nil {
				color.Red("Error closing channel", err)
			}
		}
	}
}
