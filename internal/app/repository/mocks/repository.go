package mocks

import (
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) SaveUser(code string, user *model.User) error {
	return nil
}

func (m *RepositoryMock) GetUserBy(id string) (*model.User, error) {
	user := &model.User{
		URL: "https://yandex.ru/",
	}
	return user, nil
}
