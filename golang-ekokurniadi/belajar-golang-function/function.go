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

	goodBye := getGoodBye
	fmt.Println(goodBye("bay"))

	sayHelloWithFilter("bay", spamFilter)

	// anonymous func
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("bay", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
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

// function as value
func getGoodBye(name string) string {
	return "bye " + name
}

// function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// anonymous function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("your are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}
