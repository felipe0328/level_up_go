package main

import (
	"fmt"
	"log"
	"sync"
)

// setup constants
const baristaCount = 3
const customerCount = 20
const maxOrderCount = 40

// the total amount of drinks that the bartenders have made
type coffeeShop struct {
	orderCount int
	orderLock  sync.Mutex

	orderCoffee  chan struct{}
	finishCoffee chan struct{}
	ordersSignal chan struct{}
}

// registerOrder ensures that the order made by the baristas is counted
func (p *coffeeShop) registerOrder() {
	p.orderLock.Lock() // using mutex.lock to avoid raise conditions
	defer p.orderLock.Unlock()

	p.orderCount++

	if p.orderCount == maxOrderCount {
		close(p.ordersSignal)
	}
}

// barista is the resource producer of the coffee shop
func (p *coffeeShop) barista(name string) {
	for {
		select {
		case <-p.orderCoffee:
			p.registerOrder()
			log.Printf("%s makes a coffee.\n", name)
			p.finishCoffee <- struct{}{}
		case <-p.ordersSignal:
			log.Printf("%s ends shift.\n", name)
			return
		}
	}
}

// customer is the resource consumer of the coffee shop
func (p *coffeeShop) customer(name string) {
	for {
		select {
		case p.orderCoffee <- struct{}{}:
			log.Printf("%s orders a coffee!", name)
			<-p.finishCoffee
			log.Printf("%s enjoys a coffee!\n", name)
		case <-p.ordersSignal:
			log.Printf("%s leaves coffee!\n", name)
			return
		}
	}
}

func main() {
	log.Println("Welcome to the Level Up Go coffee shop!")
	orderCoffee := make(chan struct{}, baristaCount)
	finishCoffee := make(chan struct{}, baristaCount)
	ordersSignal := make(chan struct{})
	p := coffeeShop{
		orderCoffee:  orderCoffee,
		finishCoffee: finishCoffee,
		ordersSignal: ordersSignal,
	}
	for i := 0; i < baristaCount; i++ {
		go p.barista(fmt.Sprint("Barista-", i))
	}
	for i := 0; i < customerCount; i++ {
		go p.customer(fmt.Sprint("Customer-", i))
	}

	<-ordersSignal

	log.Println("The Level Up Go coffee shop has closed! Bye!")
}
