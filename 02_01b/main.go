package main

import (
	"flag"
	"log"
	"sync"
)

type chanObject struct {
	id      int
	message string
}

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, message string) {
	var wg sync.WaitGroup
	ch := make(chan chanObject)

	wg.Add(n)

	for i := 0; i < n; i++ {
		go repeatConcurrently(ch, &wg)
		ch <- chanObject{id: i, message: message}
	}

	wg.Wait()
	close(ch)
}

func repeatConcurrently(ch chan chanObject, wg *sync.WaitGroup) {
	m := <-ch
	log.Printf("[G%d]:%s\n", m.id, m.message)
	wg.Done()
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()
	for _, m := range messages {
		log.Println(m)
		repeat(int(*factor), m)
	}
}
