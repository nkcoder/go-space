package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 20
var arrivalRate = 10
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("--------------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarbersDonChan:  doneChan,
		ClientsChan:     clientChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	shop.AddBarber("Frank")

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		// creates a channel that will receive a message after the timeOpen duration, blocking
		<-time.After(timeOpen)
		shopClosing <- true
		shop.CloseShopForDay()
		closed <- true
	}()

	i := 1

	go func() {
		for {
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Microsecond * time.Duration(randomMilliseconds)):
				shop.AddClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	<-closed
}
