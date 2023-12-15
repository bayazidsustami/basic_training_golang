package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	person := Person{"bays"}
	sayHello(person)

	animal := Animal{"kucing"}
	sayHello(animal)

	ups := ups()
	fmt.Println(ups)

	anything := returnAnything()
	fmt.Println(anything)
}

// interface kosong / any
func ups() interface{} {
	// return "anything"
	return 0
}

func returnAnything() any {
	// return 0
	return "test"
}
