package main

import "fmt"

func main() {
	var fruits = []string{"apple", "banana", "melon", "manggo"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println("-------------before----------")

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	baFruits[0] = "pinnaple"

	fmt.Println("-------------after----------")
	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	fmt.Println("some operation fuction of slice")
	var lengthOfFruits = len(fruits)
	fmt.Println(lengthOfFruits)

	//cap() for calculate max capacity of slice
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))

	var appendFruits = append(fruits, "pir")
	fmt.Println(appendFruits)

	// function copy
	destination := make([]string, 3)
	source := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(destination, source)

	fmt.Println(destination)
	fmt.Println(source)
	fmt.Println(n)

	//another example of copy slice
	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinneaple"}
	nCopy := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(nCopy)

	//combine slice and map
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	fmt.Println(data[0:3])

}
