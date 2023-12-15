package main

import "fmt"

func main() {
	name := "bay"

	switch name {
	case "bay":
		fmt.Println("hello bay")
	case "yazid":
		fmt.Println("hello yazid")
	default:
		fmt.Println("unknown")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	nameLength := len(name)
	switch {
	case nameLength > 10:
		fmt.Println("nama terlalu panjang")
	case nameLength > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama benar")
	}
}
