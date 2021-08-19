package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/spinel/go-musthave-shortener-tpl/internal/app/handler"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/repository"
)

func main() {
	repo, err := repository.New()
	if err != nil {
		errors.Wrap(err, "store")
	}
	r := mux.NewRouter()
	r.HandleFunc("/", handler.CreateUserHandler(repo))
	r.HandleFunc("/{id:[0-9a-z]+}", handler.GetUserHandler(repo))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
