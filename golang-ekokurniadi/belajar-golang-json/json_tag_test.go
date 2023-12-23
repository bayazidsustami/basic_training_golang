package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTagEncode(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "MB Pro 1",
		Price:    24000000,
		ImageUrl: "http://test.com/contoh.png",
	}

	LogJson(product)
}

func TestJsonTagDecode(t *testing.T) {
	jsonStr := `{"id":"P0001","name":"MB Pro 1","price":24000000,"image_url":"http://test.com/contoh.png"}`
	jsonBytes := []byte(jsonStr)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
