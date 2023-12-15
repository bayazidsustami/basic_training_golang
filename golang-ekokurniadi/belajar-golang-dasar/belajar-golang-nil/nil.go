package main

import "fmt"

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println("data", data)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
