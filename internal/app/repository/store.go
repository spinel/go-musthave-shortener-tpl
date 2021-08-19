package repository

import "github.com/spinel/go-musthave-shortener-tpl/internal/app/repository/web"

type Store struct {
	User Repository
}

func New() (*Store, error) {
	var store Store
	store.User = web.NewUserRepo()
	return &store, nil
}
