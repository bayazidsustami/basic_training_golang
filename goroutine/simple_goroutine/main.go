package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2) //define max core for execution program

	go print(5, "Hallo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
