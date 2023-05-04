package repository

import "golang-latihan-unit-test/entity"

type CategoryRepository interface {
	FindById(Id string) *entity.Category
}