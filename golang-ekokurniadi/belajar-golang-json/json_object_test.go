package belajargolangjson

import "testing"

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
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
