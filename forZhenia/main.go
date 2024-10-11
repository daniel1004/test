package main

import (
	"flag"
	"log"
)

func main() {
	//t: mustToken()
}
func mustToken() string {
	token := flag.String("token-bot-token", "", "token to access to telegram")
	flag.Parse()
	if *token == "" {
		log.Fatal("token is required")
	}
	return *token
}
