package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`e([a-z]o)`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("e5o"))

	fmt.Println(regex.FindAllString("eko edo eto e5o eKo", 10))
}
