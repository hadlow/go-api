package main

import (
	"log"
	"net/http"
)

// External
import (
	bolt "go.etcd.io/bbolt"
	"github.com/gorilla/mux"
)

func all(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	// db, err := bolt.Open("main.db", 0600, nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	posts := api.PathPrefix("/posts").Subrouter()

    posts.HandleFunc("/", all).Methods(http.MethodGet)
	posts.HandleFunc("/{id}", get).Methods(http.MethodGet)
	posts.HandleFunc("/", post).Methods(http.MethodPost)
	posts.HandleFunc("/{id}", put).Methods(http.MethodPut)

    log.Fatal(http.ListenAndServe(":8080", r))
}