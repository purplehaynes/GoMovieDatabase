package main

import (
	"GoMovieDB/handlers"
	"log"
)

func main() {
	server := handlers.NewServer()

	log.Fatal(server.ListenAndServe())
}