package repository

import (
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
)

type Repository interface {
	GetUserBy(string) (*model.User, error)
	SaveUser(string, *model.User) error
	IncludesCode(string) bool
}
