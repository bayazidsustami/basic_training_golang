package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "bayazid",
		MiddleName: "sustami",
		LastName:   "M N",
		Age:        26,
		Hobbies:    []string{"badmin", "gaming", "sleep"},
	}
	LogJson(customer)
}

func TestJsonArrayDecode(t *testing.T) {
	jsonStr := `{"FirstName":"bayazid","MiddleName":"sustami","LastName":"M N","Age":26,"Hobbies":["badmin","gaming","sleep"]}`
	jsonBytes := []byte(jsonStr)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestJsonArrayComplexEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "bayazid",
		MiddleName: "sustami",
		LastName:   "M N",
		Addresses: []Address{
			{
				Street:     "Jl jalan",
				Country:    "INA",
				PostalCode: 90233,
			},
			{
				Street:     "Jl jalanan",
				Country:    "MY",
				PostalCode: 902324,
			},
			{
				Street:     "Jl aja",
				Country:    "SGP",
				PostalCode: 902332,
			},
		},
	}
	LogJson(customer)
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonStr := `{"FirstName":"bayazid","MiddleName":"sustami","LastName":"M N","Age":0,"Hobbies":null,"Addresses":[{"Street":"Jl jalan","Country":"INA","PostalCode":90233},{"Street":"Jl jalanan","Country":"MY","PostalCode":902324},{"Street":"Jl aja","Country":"SGP","PostalCode":902332}]}`
	jsonBytes := []byte(jsonStr)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
