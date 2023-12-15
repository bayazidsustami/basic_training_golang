package main

import "fmt"

func main() {
	name := "bay"

	if name == "bay" {
		fmt.Println("hello bay")
	} else if name == "yazid" {
		fmt.Println("hello yazid")
	} else {
		fmt.Println("unknown")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
