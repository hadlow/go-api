package main

import (
	"log"
	"net/http"
)

// External
import (
	// bolt "go.etcd.io/bbolt"
	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
	db database
}

func (a *App) Initialize(dbPath string) {
	a.Router = mux.NewRouter()

	// Get the database from whatever file is in the config
	database, close, err := database.NewDatabase(config.Database)

	database.SetBucket("main")

	if err != nil {
		log.Fatal("Error opening database")
	}
}

func (a *App) Run(addr string) {
	api := a.Router.PathPrefix("/api").Subrouter()

    api.HandleFunc("/posts", all).Methods(http.MethodGet)
	api.HandleFunc("/posts/{id}", get).Methods(http.MethodGet)
	api.HandleFunc("/posts", post).Methods(http.MethodPost)
	api.HandleFunc("/posts/{id}", put).Methods(http.MethodPut)
	api.HandleFunc("/posts/{id}", delete).Methods(http.MethodDelete)

    log.Fatal(http.ListenAndServe(":8080", a.Router))
}