package belajargolanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, float32(20.00), ExperimentalMin(float32(20.00), float32(23.00)))
	assert.Equal(t, float64(200.00), ExperimentalMin(float64(200.00), float64(230.00)))
	assert.Equal(t, Age(200), ExperimentalMin(Age(200), Age(230)))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Bay",
	}
	second := map[string]string{
		"Name": "Bay",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlice(t *testing.T) {
	first := []string{"bay"}
	second := []string{"bay"}

	assert.True(t, slices.Equal(first, second))
}
