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
        w.Write([]byte("OK"))
    })
	r.Mount("/books", BookRoutes())
    http.ListenAndServe(":3001", r)
}

func BookRoutes() chi.Router {
    r := chi.NewRouter()
    bookHandler := BookHandler{}
    r.Get("/", bookHandler.ListBooks)
    r.Post("/", bookHandler.CreateBook)
    r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)
    return r
}