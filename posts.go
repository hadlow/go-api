package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io"
	"time"
)

type Post struct {
	ID int     `json:"id"`
	Title string  `json:"title"`
	Content string `json:"content"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
	DeletedOn string  `json:"-"`
}

type Posts []*Post

func (p *Posts) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetPosts() Posts {
	return postsList
}

var postsList = []*Post {
	&Post{
		ID: 1,
		Title: "Latte",
		Content: "Frothy milky coffee",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}

type PostsApi struct {
	l *log.Logger
}

func NewPostsApi(l *log.Logger) *PostsApi {
	return &PostsApi{l}
}

func (p *PostsApi) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getPosts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *PostsApi) getPosts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Posts")

	lp := GetPosts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}