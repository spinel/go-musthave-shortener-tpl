package mocks

import (
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
	"github.com/stretchr/testify/mock"
)

// RepositoryMock is a struct for mock tests
type RepositoryMock struct {
	mock.Mock
}

// SaveEntity is a fake method for save entity.
func (m *RepositoryMock) SaveEntity(code string, entity model.Entity) error {
	return nil
}

// GetEntityBy is a fake get method to get entity by id.
func (m *RepositoryMock) GetEntityBy(id string) (model.Entity, error) {
	entity := model.Entity{
		URL: "https://yandex.ru/",
	}
	return entity, nil
}
