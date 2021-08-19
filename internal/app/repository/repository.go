package repository

import (
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
	"github.com/stretchr/testify/mock"
)

type Repository interface {
	GetUserBy(string) (string, error)
	SaveUser(*model.User) (string, error)
}

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetUserBy(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)

}
