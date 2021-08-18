package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/handler"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/storage"
)

func main() {
	s := make(storage.Storage)
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Create(s))
	r.HandleFunc("/{id:[0-9a-z]+}", handler.Get(s))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
