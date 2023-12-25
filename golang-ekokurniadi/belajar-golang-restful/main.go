package main

import (
	"belajar-golang-restful/app"
	"belajar-golang-restful/controller"
	"belajar-golang-restful/repository"
	"belajar-golang-restful/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	controller := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", controller.FindAll)
	router.GET("/api/categories/:categoryId", controller.FindById)
	router.POST("/api/categories", controller.Create)
	router.PUT("/api/categories/:categoryId", controller.Update)
	router.DELETE("/api/categories/:categoryId", controller.Delete)

}
