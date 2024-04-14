package main

import (
	"log"

	"github.com/rackerlabs/goraxauth"
)

func main() {
	opts, err := goraxauth.AuthOptionsFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	provider, err := goraxauth.AuthenticatedClient(opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Token: %s", provider.TokenID)
}
