package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("Initial balance: $%d\n", bankBalance)

	incomes := []Income{
		{Source: "Job", Amount: 600},
		{Source: "Gifts", Amount: 20},
		{Source: "Investments", Amount: 100},
		{Source: "Part time", Amount: 30},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d, Earned $%d.00, from %s\n", week, income.Amount, income.Source)

			}
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("Sum: %d", bankBalance)
}
