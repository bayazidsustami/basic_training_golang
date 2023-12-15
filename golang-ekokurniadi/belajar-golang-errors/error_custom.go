package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "bay" {
		return &notFoundError{"data tidak ditemukan"}
	}

	return nil
}

func main() {
	err := saveData("test")

	if err != nil {
		if validationError, ok := err.(*validationError); ok {
			fmt.Println("validation error :", validationError.Message)
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("not found error", notFoundError.Message)
		} else {
			fmt.Println("unknown error", err.Error())
		}
	} else {
		fmt.Println("success")
	}
}
