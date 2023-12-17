package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("bayazid", "bay"))
	fmt.Println(strings.Split("bayazid sustami", " "))
	fmt.Println(strings.ToLower("Bayazid Sustami"))
	fmt.Println(strings.ToUpper("Bayazid Sustami"))
	fmt.Println(strings.Trim("           bayazid         ", " "))
	fmt.Println(strings.ReplaceAll("Bayazid", "Bay", "*"))
}
