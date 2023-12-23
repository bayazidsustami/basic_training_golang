package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapDecode(t *testing.T) {
	jsonStr := `{"id":"P0001","name":"MB Pro 1","price":24000000,"image_url":"http://test.com/contoh.png"}`
	jsonBytes := []byte(jsonStr)

	var result map[string]any
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestJsonMapEncode(t *testing.T) {
	product := map[string]any{
		"id":    "P0001",
		"name":  "MB Pro 1",
		"price": 24000000,
	}

	LogJson(product)
}
