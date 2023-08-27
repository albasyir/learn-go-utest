package repository

import "utest/entity"

type CategoryRepository interface {
	FindById(id int) *entity.Category
}
