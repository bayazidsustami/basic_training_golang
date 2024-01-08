// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"belajar-golang-database-migration/app"
	"belajar-golang-database-migration/controller"
	"belajar-golang-database-migration/middleware"
	"belajar-golang-database-migration/repository"
	"belajar-golang-database-migration/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"net/http"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	categoryRepository := repository.NewCategoryRepository()
	db := app.NewDB()
	v := _wireValue
	validate := validator.New(v...)
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := app.NewServer(authMiddleware)
	return server
}

var (
	_wireValue = []validator.Option{}
)

// injector.go:

var categorySet = wire.NewSet(repository.NewCategoryRepository, service.NewCategoryService, controller.NewCategoryController)

var validatorSet = wire.NewSet(validator.New, wire.Value([]validator.Option{}))