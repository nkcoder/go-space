package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3                  // how many times a philosopher eats
var eatTime = 1 * time.Second   // how long it takes to eat
var thinkTime = 3 * time.Second // how long it takes to think
var sleepTime = 1 * time.Second // how long it takes to sleep

var orderMutex sync.Mutex  // a mutex to sync access to the orderFinished slice
var orderFinished []string // a slice of names of philosophers who have finished eating

func main() {
	fmt.Println("Dining Philosophers Problem")

	fmt.Println("--------------------------------")
	fmt.Println("The table is empty.")

	time.Sleep(sleepTime)

	dine()

	fmt.Println("--------------------------------")
	fmt.Println("The table is empty.")
	time.Sleep(sleepTime)
	fmt.Printf("Order finished: %s.\n", strings.Join(orderFinished, ", "))
}

func dine() {

	// Keep tract of how many philosophers are still at the table, when it reaches zero, everyone is finished eating and has left.
	atTable := &sync.WaitGroup{}
	atTable.Add(len(philosophers))

	// we want everyone to be seated before they start eating
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// each fork has a unique mutex
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], atTable, forks, seated)
	}

	atTable.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", philosopher.name)

	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s picked up the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s picked up the right fork.\n", philosopher.name)
		} else {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s picked up the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s picked up the left fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	fmt.Printf("%s is satisfied.\n", philosopher.name)
	fmt.Printf("%s left the table.\n", philosopher.name)

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()

	fmt.Println("--------------------------------")
}
