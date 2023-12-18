package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func HelloWord() {
	fmt.Println("hello world")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestCreateGoroutine(t *testing.T) {
	go HelloWord()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
