package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "jhon wick"
	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded :", encodedString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodeByte)
	fmt.Println("decoded :", decodedString)

	encodeDecode()
}

func encodeDecode() {
	var data = "ethan hunt"
	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString = string(encoded)
	fmt.Println(encodedString)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}

	var decodedString = string(decoded)
	fmt.Println(decodedString)
}
