package controller

import (
	"belajar-golang-database-migration/model/web"
	"belajar-golang-database-migration/service"
	"belajar-golang-database-migration/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}

	utils.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   categoryResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}

	utils.ReadFromRequestBody(request, &categoryUpdateRequest)

	paramCategoryId := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(paramCategoryId)
	utils.PanicErr(err)

	categoryUpdateRequest.Id = categoryId

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   categoryResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paramCategoryId := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(paramCategoryId)
	utils.PanicErr(err)

	controller.CategoryService.Delete(request.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paramCategoryId := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(paramCategoryId)
	utils.PanicErr(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   categoryResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponse := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   categoriesResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}
