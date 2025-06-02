package main

import (
	"fmt"
	"sync"
)

// 1. Update message example
var msg string
var wg sync.WaitGroup

func update(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

// go run -race .
// -race is used to execute a go program while simultaneously enabling the built-in race detector.
func updateMessage() {
	msg = "Hello, World!"

	wg.Add(2)

	var mutex sync.Mutex

	go update("Hello, Galaxy!", &mutex)
	go update("Hello, Universe!", &mutex)

	wg.Wait()

	fmt.Println(msg)
}

// 2. Calculate income example
type income struct {
	source string
	amount int
}

// processIncome handles the weekly income processing for a single income source
func processIncome(i income, balanceMutex *sync.Mutex, bankBalance *int) {
	defer wg.Done()

	for week := 1; week <= 52; week++ {
		balanceMutex.Lock()
		*bankBalance += i.amount
		balanceMutex.Unlock()

		fmt.Printf("On week %d, you earned %d from %s\n", week, i.amount, i.source)
	}
}

func calculateIncome() {
	var bankBalance int
	var balanceMutex sync.Mutex

	fmt.Println("Initial bank balance:", bankBalance)

	incomes := []income{
		{source: "Main job", amount: 1000},
		{source: "Part-time job", amount: 500},
		{source: "Investments", amount: 100},
		{source: "Gifts", amount: 10},
	}

	wg.Add(len(incomes))

	for _, i := range incomes {
		go processIncome(i, &balanceMutex, &bankBalance)
	}

	wg.Wait()

	fmt.Printf("Final bank balance: %d\n", bankBalance)
}
