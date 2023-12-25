package service

import (
	"belajar-golang-restful/model/domain"
	"belajar-golang-restful/model/web"
	repository "belajar-golang-restful/repository"
	"belajar-golang-restful/utils"
	"context"
	"database/sql"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	utils.PanicErr(err)
	defer utils.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return utils.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	utils.PanicErr(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	utils.PanicErr(err)

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return utils.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	utils.PanicErr(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	utils.PanicErr(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	utils.PanicErr(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	utils.PanicErr(err)

	return utils.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	utils.PanicErr(err)
	defer utils.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return utils.ToCategoriesResponse(categories)
}
