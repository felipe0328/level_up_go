package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	ID   string
	Name string
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	f, e := os.Open(path)

	if e != nil {
		log.Fatal("Error processing file: ", e)
	}

	defer f.Close()

	dec := json.NewDecoder(f)

	var raffle []raffleEntry

	e = dec.Decode(&raffle)

	if e != nil {
		log.Fatal("Error processing file: ", e)
	}

	return raffle
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
