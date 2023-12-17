package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"john", "paul", "george", "ringo"}
	values := []int{10, 20, 40, 15, 30}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "paul"))
	fmt.Println(slices.Index(names, "george"))
}
