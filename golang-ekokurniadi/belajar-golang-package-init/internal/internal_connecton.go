package internal

import "fmt"

// auto init ketika pertama kali diakses atau di import menggunakan blank identifier(_)
func init() {
	fmt.Println("connected")
}
