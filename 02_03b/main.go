package main

import (
	"fmt"
	"log"
)

// the number of attendees we need to serve lunch to
const (
	consumerCount = 300
)

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(consumerName string, food []chan string, doneSignal chan<- struct{}) {
	for _, ch := range food {
		fmt.Printf("%s takes %s\n", consumerName, <-ch)
	}

	doneSignal <- struct{}{}
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(food string, out chan<- string, done <-chan struct{}) {
	for {
		select {
		case out <- food:
		case <-done:
			return
		}
	}
}

func startServingFood(numberOfCosumers int) {
	courses := make([]chan string, len(foodCourses))
	doneServing := make(chan struct{})
	doneEating := make(chan struct{})

	for index, value := range foodCourses {
		ch := make(chan string)
		courses[index] = ch
		go serveLunch(value, ch, doneServing)
	}

	for i := 0; i < numberOfCosumers; i++ {
		name := fmt.Sprintf("Consumer %d", i)
		go takeLunch(name, courses, doneEating)
	}

	for i := 0; i < numberOfCosumers; i++ {
		<-doneEating
	}

	close(doneServing)
}

func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n",
		consumerCount)

	startServingFood(consumerCount)
}
