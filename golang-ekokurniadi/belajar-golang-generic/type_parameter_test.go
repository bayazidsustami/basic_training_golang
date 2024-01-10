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

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestSampleGeneric(t *testing.T) {
	var result string = Length[string]("Bayazid")
	assert.Equal(t, "Bayazid", result)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("test string", 100)
	MultipleParameter[float32, bool](23.4, true)
}
