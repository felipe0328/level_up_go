package main

import (
	"encoding/json"
	"log"
	"os"
)

// User represents a user record.
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

const path = "users.json"

// getBiggestMarket takes in the slice of users and
// returns the biggest market.
func getBiggestMarket(users []User) (string, int) {
	countryUsers := make(map[string]int)

	for _, user := range users{
		countryUsers[user.Country]++
	}

	var maxUserCountryAmount int
	var biggestUsersCountry string

	for country, amount := range countryUsers {
		if amount > maxUserCountryAmount {
			maxUserCountryAmount = amount
			biggestUsersCountry = country
		}
	}

	return biggestUsersCountry, maxUserCountryAmount
}

func main() {
	users := importData()
	country, count := getBiggestMarket(users)
	log.Printf("The biggest user market is %s with %d users.\n",
		country, count)
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []User {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []User
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
