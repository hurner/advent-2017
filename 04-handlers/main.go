package main

import (
	"log"
	"net/http"

	"github.com/rumyantseva/advent-2017/04-handlers/handlers"
)

// How to try it: go run main.go
func main() {
	log.Print("Starting the service...")
	r := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", r))
}