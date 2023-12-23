package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonStr := `{"FirstName":"bayazid","MiddleName":"sustami","LastName":"M N","Age":26}`
	jsonBytes := []byte(jsonStr)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
