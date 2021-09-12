package main

import "fmt"

func main() {
	var data = student{
		name:        "wick",
		height:      182.5,
		age:         26,
		isGraduated: false,
		hobbies:     []string{"eating", "sleeping"},
	}

	//to format numeric base biner
	fmt.Printf("%b \n", data.age)
	//to format unicode
	fmt.Printf("%c \n", 1235)
	//to format numeric
	fmt.Printf("%d \n", data.age)
	//to format decimal into standart numeric anotation
	fmt.Printf("%e \n", data.height)
	//to format decimal and can adjust width of number
	fmt.Printf("%f \n", data.height)
	fmt.Printf("%.2f \n", data.height)
	fmt.Printf("%.5f \n", data.height)
	fmt.Printf("%g \n", data.height)
	//to format oktal
	fmt.Printf("%o \n", data.age)
	//to format pointer
	fmt.Printf("%p \n", &data.name)
	//to format escape
	fmt.Printf("%q\n", `" name \ height "`)
	//to format boolean
	fmt.Printf("%t \n", data.isGraduated)
	//to get type of variable
	fmt.Printf("%T \n", data.name)
	//to format everything type
	fmt.Printf("%v \n", data)
	fmt.Printf("%+v \n", data)
	fmt.Printf("%+#v \n", data)
	//to format string hexadecimal
	fmt.Printf("%x\n", data.age)

}

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}
