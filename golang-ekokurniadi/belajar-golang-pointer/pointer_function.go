package main

import "fmt"

type Address struct {
	City, Region, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Mks", "Sulsel", ""}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}
