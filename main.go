package main

import (
	"log"
	"net/http"
	"os"
)

const dbfilename = "game.db.json"

func main() {
	db, err := os.OpenFile(dbfilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbfilename, err)
	}

	store := &FileSystemPlayerStore{db}
	server := NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
