package repository

import (
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/repository/web"
)

// Store main struct.
type Store struct {
	MemoryDB map[string]model.Entity
	Entity   Repository
}

// New store builder.
func New() (*Store, error) {
	db := make(map[string]model.Entity)
	store := Store{
		MemoryDB: db,
		Entity:   web.NewEntityRepo(db),
	}
	return &store, nil
}
