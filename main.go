package main

import (
	"log"
	"net/http"

	"github.com/omekov/exclusive/handlers"
)

func main() {
	router := handlers.StartServer()
	log.Fatal(http.ListenAndServe(":5555", router))
}
