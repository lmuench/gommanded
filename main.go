package main

import (
	"log"
	"net/http"

	"github.com/lmuench/gommanded/web"
)

const port = ":5000"

func main() {
	router := web.Router()
	log.Println("Server listening on", port)
	log.Fatal(http.ListenAndServe(port, router))
}
