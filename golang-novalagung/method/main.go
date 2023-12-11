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

	s1.changeName1("jason bourne")
	fmt.Println("s1 after change name : ", s1.name)

	s1.changeName2("ethan hunt")
	fmt.Println("s2 after change name 2 : ", s1.name)

	var s2 = &student{"ethan", 23}
	s2.sayHello()
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

func (s student) changeName1(name string) {
	fmt.Println("-------> on ChangeName1, name change to ", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("-------> on ChangeName2, name change to ", name)
	s.name = name
}
