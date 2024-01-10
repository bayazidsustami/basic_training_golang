package belajargolanggeneric

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bags Bag[T]) {
	for _, v := range bags {
		fmt.Println(v)
	}
}

func TestPrintBagString(t *testing.T) {
	names := Bag[string]{"eiger", "arei", "bodypack"}
	PrintBag(names)

	numbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(numbers)
}
