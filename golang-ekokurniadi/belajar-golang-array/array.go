package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "bay"
	names[1] = "baya"
	names[2] = "bayazid"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		96,
		97,
		98,
	}
	fmt.Println(values)
	fmt.Println(len(values))

	values[1] = 20
	fmt.Println(values)

	var newValues = [...]int{1, 2, 3}
	fmt.Println(len(newValues))

}
