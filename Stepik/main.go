package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	input, err := os.Open("transaction_logs.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	var email, number string
	reader := csv.NewReader(input)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range data {
		email = row[0]
		number = row[1]
	}

}
