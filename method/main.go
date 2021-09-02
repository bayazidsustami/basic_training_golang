package main

import (
	"fmt"
	"strings"
)

func main() {
	sayHello()

	var s1 = student{"jhon wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("Nama panggilan", name)
}

type student struct {
	name  string
	grade int
}

//function
func sayHello() {
	fmt.Println("hello")
}

//method
func (s student) sayHello() {
	fmt.Println("hello", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}
