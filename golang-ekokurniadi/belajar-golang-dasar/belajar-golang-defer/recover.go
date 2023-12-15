package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("end app")

	message := recover()
	fmt.Println("terjadi error", message)
}

func runApp(isError bool) {
	defer endApp()
	if isError {
		panic("ERROR")
	}
}
