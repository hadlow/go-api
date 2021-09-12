package main

import (
	"net/http"
)

// External
import (
	"github.com/gorilla/mux"
)

// External
import (
	// bolt "go.etcd.io/bbolt"
	// "github.com/gorilla/mux"
)

type Post struct {
	id int `json:"id"`
	title string `json:"title"`
	body string `json:"body"`
}

func PostsApi(api *mux.Router) {
	api.HandleFunc("/posts", AllPosts).Methods(http.MethodGet)
	api.HandleFunc("/posts/{id}", GetPosts).Methods(http.MethodGet)
	api.HandleFunc("/posts", PostPosts).Methods(http.MethodPost)
	api.HandleFunc("/posts/{id}", PutPosts).Methods(http.MethodPut)
	api.HandleFunc("/posts/{id}", DeletePosts).Methods(http.MethodDelete)
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// vars["id"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func PostPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func PutPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func DeletePosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}