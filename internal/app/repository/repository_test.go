package repository

import (
	"errors"
	"testing"

	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/repository/mocks"
)

// TestGetEntityBy 
func TestGetEntityBy(t *testing.T) {
	repoMock := new(mocks.RepositoryMock)

	repoMock.On("SaveEntity", "testtest", &model.Entity{URL: "new"}).Return(nil)
	repoMock.AssertExpectations(t)
	repoMock.On("SaveEntity", "", &model.Entity{URL: "new"}).Return(nil)

	repoMock.On("GetEntityBy", "testtest").Return("name")
	repoMock.On("GetEntityBy", "").Return("", errors.New("not found"))
}
