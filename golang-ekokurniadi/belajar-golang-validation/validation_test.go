package belajargolangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := ""

	err := validate.Var(user, "required")

	assert.NotNil(t, err)
}

func TestValidationTwoVariable(t *testing.T) {
	validate := validator.New()
	password := "password"
	confirmPassword := "salah"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	assert.NotNil(t, err)
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "12345"

	err := validate.Var(user, "required,numeric")
	assert.Nil(t, err)
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "99999999999999"

	err := validate.Var(user, "required,numeric,min=5,max=10")
	assert.NotNil(t, err)
}

func TestValidationStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "eko@example.com",
		Password: "eko1234",
	}

	err := validate.Struct(loginRequest)
	assert.Nil(t, err)
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "eko",
		Password: "eko",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestValidationStructCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	request := RegisterUser{
		Username:        "eko@example.com",
		Password:        "eko1234",
		ConfirmPassword: "eko1234",
	}

	err := validate.Struct(request)
	assert.Nil(t, err)
}

func TestValidationNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	validate := validator.New()
	request := User{
		Id:   "01",
		Name: "test",
		Address: Address{
			City:    "test",
			Country: "test",
		},
	}

	err := validate.Struct(request)
	assert.Nil(t, err)
}
