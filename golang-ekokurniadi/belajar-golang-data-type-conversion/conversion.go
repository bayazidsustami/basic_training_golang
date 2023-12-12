package main

import "fmt"

func main() {
	var value32 int32 = 32768
	var value64 int64 = int64(value32) // konversi dari type int32 ke int64
	var value16 int16 = int16(value32) // konversi dari type int32 ke int16

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16)

	name := "bayazid sustami"
	firstCharInByte := name[0]                      //char yang diambil berupa byte uint8
	var firstCharInString = string(firstCharInByte) // harus dikonversi ke string

	fmt.Println(firstCharInByte)
	fmt.Println(firstCharInString)
}
