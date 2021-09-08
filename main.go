package main

import (
	"log"
)

import (
	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("main.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}