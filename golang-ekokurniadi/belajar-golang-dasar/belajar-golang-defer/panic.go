package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("end app")
}

func runApp(isError bool) {
	defer endApp()
	if isError {
		panic("ERROR")
	}
}
