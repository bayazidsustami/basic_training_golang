package belajargolanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin[int](100, 200))
	assert.Equal(t, float32(20.00), FindMin[float32](20.00, 23.00))
	assert.Equal(t, float64(200.00), FindMin[float64](200.00, 230.00))
}

func FindFirst[T []E, E any](data T) E {
	first := data[0]
	return first
}

func TestFindFirst(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	first := FindFirst[[]int, int](numbers)
	assert.Equal(t, 1, first)
}
