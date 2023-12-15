package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "bay",
		"address": "Mks",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Go-Lang"
	book["author"] = "bay"
	book["wrong"] = "ups"

	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)
}
