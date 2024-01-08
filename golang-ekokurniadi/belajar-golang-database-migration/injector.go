//go:build wireinject
// +build wireinject

package main

import (
	"belajar-golang-database-migration/app"
	"belajar-golang-database-migration/controller"
	"belajar-golang-database-migration/middleware"
	"belajar-golang-database-migration/repository"
	"belajar-golang-database-migration/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

var validatorSet = wire.NewSet(
	validator.New,
	wire.Value([]validator.Option{}),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validatorSet,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		app.NewServer,
	)
	return nil
}
