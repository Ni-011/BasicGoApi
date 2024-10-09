package main

import (
	"log"
)

func main() {
	// create the db
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	// init tables
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	// run server
	server := NewServer(":8000", store)
	server.Run()
}