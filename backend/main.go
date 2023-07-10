package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	r.Post("/add", func(w http.ResponseWriter, r *http.Request) {

	})
	r.Get("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":3000", r)
}