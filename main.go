package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/JaraKramar/golang-chi-crud-api/books"
)

func main() {
	r := setupServer()
	// Listen on port 3001 check if is available 
	// if not use another port
	http.ListenAndServe(":3001", r)
}

func setupServer() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Mount("/books", BookRoutes())
	return r
}

func BookRoutes() chi.Router {
	r := chi.NewRouter()
	bookHandler := books.BookHandler{Storage: books.BookStore{}}
	r.Get("/", bookHandler.ListBooks)
	r.Post("/", bookHandler.CreateBook)
	r.Get("/{id}", bookHandler.GetBooks)
	r.Put("/{id}", bookHandler.UpdateBook)
	r.Delete("/{id}", bookHandler.DeleteBook)
	return r
}