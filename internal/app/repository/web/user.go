package web

import (
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
)

// UserRepo ...
type UserRepo struct {
	Memory map[string]*model.User
}

// NewUserRepo ...
func NewUserRepo(db map[string]*model.User) *UserRepo {
	var repo UserRepo
	repo.Memory = db
	return &repo
}

func (repo *UserRepo) GetUserBy(id string) (*model.User, error) {
	user := repo.Memory[id]
	return user, nil
}

func (repo *UserRepo) SaveUser(code string, user *model.User) error {
	repo.Memory[code] = user
	return nil
}

func (repo *UserRepo) IncludesCode(code string) bool {
	_, result := repo.Memory[code]
	return result
}
