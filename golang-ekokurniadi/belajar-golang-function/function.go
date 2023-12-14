package main

import "fmt"

func main() {
	sayHello()

	sayHelloTo("bay", "yazid")

	fmt.Println(getHello("bay"))

	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	fmt.Println(getNumbers())

}

func sayHello() {
	fmt.Println("Hellow")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hellow", firstName, lastName)
}

func getHello(name string) string {
	return "hello " + name
}

// multiple return values
func getFullName() (string, string, string) {
	return "bay", "yazid", "sustami"
}

// named return values
func getNumbers() (numb1, numb2, numb3 int) {
	numb1 = 1
	numb2 = 2
	numb3 = 3
	return numb1, numb2, numb3
}
