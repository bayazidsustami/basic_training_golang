package main

import (
	"errors"
	"fmt"
)

func division(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("tidak dapat dibagi 0")
	} else {
		return value / divider, nil
	}
}

func main() {
	result, err := division(2, 0)
	if err == nil {
		fmt.Println("result :", result)
	} else {
		fmt.Println("error", err.Error())
	}

}
