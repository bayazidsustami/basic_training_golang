package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(value any) {
	byte, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestEncodeJson(t *testing.T) {
	LogJson("eko")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"bay", "bayazid", "yazid"})
}
