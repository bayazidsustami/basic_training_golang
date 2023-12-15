package main

import "fmt"

func random() interface{} {
	return "Ok"
}

func main() {
	result := random()
	resultAsString := result.(string) // convert to string
	fmt.Println(resultAsString)

	// resultAsInt := result.(int) //panic - can't convert because random() return string
	// fmt.Println(resultAsInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
