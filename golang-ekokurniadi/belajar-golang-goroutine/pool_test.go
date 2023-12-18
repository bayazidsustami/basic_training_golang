package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} { // override value default
			return "New"
		},
	}

	pool.Put("bay")
	pool.Put("yazid")
	pool.Put("bayazid")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("data", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
}
