package belajargolanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	assert.Equal(t, float32(20.00), Min[float32](20.00, 23.00))
	assert.Equal(t, float64(200.00), Min[float64](200.00, 230.00))
}
