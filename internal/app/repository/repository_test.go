package repository

import (
	"errors"
	"testing"

	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/repository/mocks"
)

func TestGetUserBy(t *testing.T) {
	repoMock := new(mocks.RepositoryMock)

	repoMock.On("SaveUser", "testtest", &model.User{URL: "new"}).Return(nil)
	repoMock.AssertExpectations(t)
	repoMock.On("SaveUser", "", &model.User{URL: "new"}).Return(nil)

	repoMock.On("GetUserBy", "testtest").Return("name")
	repoMock.On("GetUserBy", "").Return("", errors.New("not found"))
}
