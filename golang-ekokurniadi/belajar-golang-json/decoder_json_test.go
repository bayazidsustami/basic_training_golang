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
