package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("bay")
	data.PushBack("yazid")
	data.PushBack("bayazid")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	data.PushFront("baya")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
