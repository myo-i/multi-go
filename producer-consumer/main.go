package producerconsumer

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

func main() {

}
