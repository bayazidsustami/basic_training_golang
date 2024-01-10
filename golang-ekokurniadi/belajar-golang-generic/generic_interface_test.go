package belajargolanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}
	result := ChangeValue[string](&myData, "bay")

	assert.Equal(t, "bay", result)
}
