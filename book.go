// Handlers for the router yet to be implemented
package main

import "net/http"

type BookHandler struct {
}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request)  {}
func (b BookHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {w.Write()}
func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}