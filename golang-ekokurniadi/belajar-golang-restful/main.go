package main

import (
	"belajar-golang-restful/app"
	"belajar-golang-restful/controller"
	"belajar-golang-restful/middleware"
	"belajar-golang-restful/repository"
	"belajar-golang-restful/service"
	"belajar-golang-restful/utils"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	controller := controller.NewCategoryController(categoryService)

	router := app.NewRouter(controller)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	utils.PanicErr(err)

}
