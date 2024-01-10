package belajargolanggeneric

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSampleGeneric(t *testing.T) {
	var result string = Length[string]("Bayazid")
	assert.Equal(t, "Bayazid", result)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}
