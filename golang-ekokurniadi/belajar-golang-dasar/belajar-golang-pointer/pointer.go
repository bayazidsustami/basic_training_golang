package main

import "fmt"

func main() {
	passByValue()
	fmt.Println("===========")
	passByRefference()
	fmt.Println("===========")
	asteriskOperator()
	fmt.Println("===========")
	newOperator()
}

type Address struct {
	City, Region, Country string
}

func passByValue() {
	address1 := Address{"Makassar", "Sulsel", "Indonesia"}
	address2 := address1

	address2.City = "Maros"

	fmt.Println(address1) // data tidak berubah
	fmt.Println(address2)
}

func passByRefference() {
	address1 := Address{"Makassar", "Sulsel", "Indonesia"}
	address2 := &address1 // operator & merujuk ke alamat memory/refference (pointer)

	address2.City = "Maros"

	fmt.Println(address1)
	fmt.Println(address2)
}

func asteriskOperator() {
	address1 := Address{"Makassar", "Sulsel", "Indonesia"}
	address2 := &address1

	address2.City = "Maros"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi bandung

	*address2 = Address{"Mamuju", "Sulbar", "Indonesia"}
	fmt.Println(address1) // smua berubah
	fmt.Println(address2)
}

func newOperator() {
	address := new(Address)
	address2 := address

	address2.City = "Makassar"

	fmt.Println(address) // ikut berubah
	fmt.Println(address2)
}
