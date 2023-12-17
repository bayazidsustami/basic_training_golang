package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("error", err.Error())
	}

	resultInt, err := strconv.Atoi("890")
	if err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("error", err.Error())
	}

	binary := strconv.FormatInt(10, 2)
	fmt.Println(binary)

	var intToString string = strconv.Itoa(200)
	fmt.Println(intToString)
}
