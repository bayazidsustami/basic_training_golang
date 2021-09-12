package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	defer catch()
	var input string
	fmt.Println("Type some number : ")
	fmt.Scanln(&input)

	/*var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}*/

	if valid, err := validate(input); valid {
		fmt.Println("hay", input)
	} else {
		panic(err.Error())
	}

	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message:", r)
				} else {
					fmt.Println("Application perfect")
				}
			}()
			panic("some error happen")
		}()
	}

}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
