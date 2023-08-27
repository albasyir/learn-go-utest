package service

import (
	"errors"
	"utest/entity"
	"utest/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) FindById(id int) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("Category not found")
	}

	return category, nil
}
