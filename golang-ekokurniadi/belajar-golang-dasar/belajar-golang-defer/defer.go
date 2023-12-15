package main

import "fmt"

func main() {
	runApplication()
}

func runApplication() {
	defer logging()
	fmt.Println("run application")
}

func logging() {
	fmt.Println("selesai memanggil function")
}
