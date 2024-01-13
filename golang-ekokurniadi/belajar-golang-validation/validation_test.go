package belajargolangvalidation

import (
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
