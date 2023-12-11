package main

import (
	f "fmt"
	"visibility-modifier/library"
)

func main() {
	library.SayHello("jhon")

	var s1 = library.Student{Name: "ethan", Grade: 22}
	f.Println("name :", s1.Name)
	f.Println("grade :", s1.Grade)

	sayHello(s1.Name)
}
