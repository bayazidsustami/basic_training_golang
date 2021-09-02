package main

import (
	"fmt"
	"visibility-modifier/library"
)

func main() {
	library.SayHello("jhon")

	var s1 = library.Student{Name: "ethan", Grade: 22}
	fmt.Println("name :", s1.Name)
	fmt.Println("grade :", s1.Grade)
}
