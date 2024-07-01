package main

import (
	"context"
	"log"

	"github.com/rackerlabs/goraxauth"
)

func main() {
	ctx := context.TODO()

	opts, err := goraxauth.AuthOptionsFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	provider, err := goraxauth.AuthenticatedClient(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Token: %s", provider.TokenID)
}
