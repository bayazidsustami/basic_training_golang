package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
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
		fmt.Println(err.Error())
	}

}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}
