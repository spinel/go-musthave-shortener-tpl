package repository

import "github.com/spinel/go-musthave-shortener-tpl/internal/app/model"

type Repository interface {
	GetEntityBy(id string) (model.Entity, error)
	SaveEntity(id string, entity model.Entity) error
	IncludesCode(id string) bool
}
