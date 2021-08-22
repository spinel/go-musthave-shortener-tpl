package repository

import (
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/repository/web"
)

type Store struct {
	MemoryDB map[string]*model.User
	User     Repository
}

func New() (*Store, error) {
	db := make(map[string]*model.User)
	var store Store
	store.MemoryDB = db
	store.User = web.NewUserRepo(db)
	return &store, nil
}
