package main

import (
	"belajar-golang-package/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("bay")
	fmt.Println(result)

	fmt.Println(helper.ApplicationName)
}
