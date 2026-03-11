package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=80"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/users", func(ctx echo.Context) error {
		u := new(User)
		if err := ctx.Bind(u); err != nil {
			return err
		}
		if err := ctx.Validate(u); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, true)
	})

	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObj, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObj {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must greather than %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must lower than %s", err.Field(), err.Param())
				}
				break
			}
		}

		ctx.Logger().Error(report)
		ctx.JSON(report.Code, report)
	}

	e.Logger.Fatal(e.Start(":9000"))
}
