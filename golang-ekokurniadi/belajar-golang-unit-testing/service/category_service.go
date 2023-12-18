package service

import (
	"belajar-golang-unit-testing/entity"
	"belajar-golang-unit-testing/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return category, errors.New("not found category")
	} else {
		return category, nil
	}
}
