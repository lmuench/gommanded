package main

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/lmuench/gommanded/api"
	"github.com/lmuench/gommanded/api/account/projector"
	"github.com/lmuench/gommanded/web"
)

const port = ":5000"

func main() {
	ctx, client := connectToDatastore()
	api.Init(ctx, client)
	projector.Init(ctx, client)

	router := web.Router()
	log.Println("Server listening on", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func connectToDatastore() (context.Context, *datastore.Client) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "gommanded-development")
	if err != nil {
		log.Fatal(err)
	}
	return ctx, client
}
