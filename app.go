package main

import (
	"log"
	"net/http"
)

// External
import (
	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
	db Database
}

func (a *App) Initialize(dbPath string) {
	a.router = mux.NewRouter()

	// Get the database from whatever file is in the config
	database, close, err := NewDatabase(dbPath)

	database.SetBucket("main")

	if err != nil {
		log.Fatal("Error opening database")
	}

	defer close()
}

func (a *App) Run(addr string) {
	api := a.router.PathPrefix("/api").Subrouter()

    PostsApi(api)

    log.Fatal(http.ListenAndServe(":8080", a.router))
}