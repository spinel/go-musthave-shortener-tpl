package router

import (
	"github.com/gorilla/mux"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/handler"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/repository"
)

func Router(repo *repository.Store) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.CreateUserHandler(repo))
	r.HandleFunc("/{id:[0-9a-z]+}", handler.GetUserHandler(repo))
	return r
}
