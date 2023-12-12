package main

import "fmt"

func main() {
	type NoKTP string

	var myNoKtp NoKTP = "123456"
	fmt.Println(myNoKtp)
	fmt.Println(NoKTP("11111111"))
}
