package repository

import (
	"utest/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id int) *entity.Category {
	data := repository.Mock.Called(id)
	dataInterface := data.Get(0)

	if dataInterface == nil {
		return nil
	}

	return dataInterface.(*entity.Category)
}
