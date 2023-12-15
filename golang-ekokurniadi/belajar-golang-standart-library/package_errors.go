package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "bay" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("error validasi")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("error notFound")
		} else {
			fmt.Println("uknown error")
		}
	} else {
		fmt.Println("success")
	}
}
