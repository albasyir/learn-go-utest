package service

import (
	"testing"
	"utest/entity"
	"utest/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: &categoryRepository}

func TestCategoryService_FindByIdFail(t *testing.T) {
	categoryRepository.Mock.On("FindById", 1).Return(nil)
	data, err := categoryService.FindById(1)

	assert.Nil(t, data)
	assert.NotNil(t, err)
}

func TestCategoryService_FindByIdSuccess(t *testing.T) {
	testData := entity.Category{Id: 2, Name: "Test Category"}
	categoryRepository.Mock.On("FindById", testData.Id).Return(&testData)

	resultData, err := categoryService.FindById(testData.Id)
	assert.Nil(t, err)
	assert.NotNil(t, resultData)
	assert.Equal(t, testData.Id, resultData.Id)
	assert.Equal(t, testData.Name, resultData.Name)
}
