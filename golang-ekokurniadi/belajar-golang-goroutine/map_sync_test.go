package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(group *sync.WaitGroup, data *sync.Map, value int) {
	group.Add(1)
	defer group.Done()

	data.Store(value, value)

}

func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go AddToMap(group, data, i)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println("key ", key, " value", value)
		return true
	})
}
