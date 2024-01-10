package belajargolanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	assert.True(t, IsSame[string]("bay", "bay"))
	assert.False(t, IsSame[int](199, 100))
}
