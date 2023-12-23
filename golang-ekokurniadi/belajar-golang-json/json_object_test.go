package belajargolangjson

import "testing"

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode int
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "bayazid",
		MiddleName: "sustami",
		LastName:   "M N",
		Age:        26,
	}

	LogJson(customer)
}
