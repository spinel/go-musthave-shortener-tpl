package web

import (
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/helper"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
)

// UserRepo ...
type UserRepo struct {
	Memory map[string]*model.User
}

// NewUserRepo ...
func NewUserRepo() *UserRepo {
	var repo UserRepo
	repo.Memory = make(map[string]*model.User)
	return &repo
}

func (repo *UserRepo) GetUserBy(id string) (string, error) {
	user := repo.Memory[id]
	return user.URL, nil
}

func (repo *UserRepo) SaveUser(user *model.User) (string, error) {
	code, err := helper.NewGeneratedString()
	if err != nil {
		return "", err
	}
	codeString := string(code)
	repo.Memory[codeString] = user
	return codeString, nil
}
