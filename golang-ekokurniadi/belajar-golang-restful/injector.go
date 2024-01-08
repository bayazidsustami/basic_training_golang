//go:build wireinject
// +build wireinject

package main

import (
	"belajar-golang-restful/app"
	"belajar-golang-restful/controller"
	"belajar-golang-restful/middleware"
	"belajar-golang-restful/repository"
	"belajar-golang-restful/service"
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
