package main

import "fmt"

func main() {
	a := 20
	b := 10

	fmt.Println(a + b) // jumlah
	fmt.Println(a - b) // kurang
	fmt.Println(b - a) // kurang
	fmt.Println(a * b) // kali
	fmt.Println(a / b) // bagi
	fmt.Println(a % b) // modulo

	a += 10
	fmt.Println(a)

	a++
	fmt.Println(a)
}
