package belajargolanggeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[F any, S any] struct {
	First  F
	Second S
}

func (d *Data[_, _]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[F, S]) ChangeFirst(first F) F {
	d.First = first
	return d.First
}

func TestGenericStruct(t *testing.T) {
	data := Data[int, string]{
		First:  10,
		Second: "sepuluh",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[int, string]{
		First:  10,
		Second: "bay",
	}

	assert.Equal(t, "Hello bay", data.SayHello("bay"))
	assert.Equal(t, 20, data.ChangeFirst(20))
}
