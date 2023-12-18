package repository

import entity "belajar-golang-unit-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
