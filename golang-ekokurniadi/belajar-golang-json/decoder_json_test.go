package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoderStream(t *testing.T) {
	reader, err := os.Open("customer.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	err = decoder.Decode(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestEncoderStream(t *testing.T) {
	writer, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "bayazid",
		MiddleName: "sustami",
		LastName:   "M N",
		Age:        26,
	}

	err := encoder.Encode(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
