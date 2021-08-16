package main

import "fmt"

func main() {
	var myName string = "bay" //initialize variable
	myLastName := "wick"

	var multiString1, multiString2, multiString3 string = "boy", "bay", "buy"

	fmt.Printf("Halo %s %s", myName, myLastName+"\n")
	fmt.Printf("name multiple %s %s %s \n", multiString1, multiString2, multiString3)

	name := new(string)
	fmt.Println(name)
	fmt.Println(*name) //referece memory address

	_ = "trash bucket" //declare unused variable with underscore

	var positiveNumber uint8 = 89
	var negativeNumber = -129304932

	fmt.Println("------------------------//type-data//---------------")
	fmt.Printf("bilangan positif : %d\n", positiveNumber)
	fmt.Printf("bilangan negatif : %d\n", negativeNumber)

	var decimalNumber = 2.30
	fmt.Printf("bilangan desimal : %f\n", decimalNumber)
	fmt.Printf("bilangan desimal : %.3f\n", decimalNumber)

	var isExist bool = true
	fmt.Printf("exist ? %t \n", isExist)

	var message string = `
	Hello my name is bay
	i'm using backticks,
	not escaping character
	`
	fmt.Print(message)

	const pi float32 = 3.14 //const is constant value, can't change
	fmt.Printf("nilai pi : %.2f \n", pi)

	fmt.Println("------------------------//type-data//---------------")
	fmt.Println("------------------------//operator-aritmatika//---------------")

	var someValue = (((2+6)%3)*4 - 2) / 3
	var modulus = 8 % 3
	fmt.Printf("value after operation : %d \n", someValue)
	fmt.Printf("modulus : %d \n", modulus)

	fmt.Println("------------------------//operator-aritmatika//---------------")
	fmt.Println("------------------------//control-flow//---------------")
	var point = 5

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 && point < 10 {
		fmt.Println("lulus")
	} else if point >= 4 && point <= 5 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus, nilai anda %d\n", point)
	}

	fmt.Println("---------------------------------------")

	//if-else with temporary variable
	var anotherPoint = 8840.0

	if percent := anotherPoint / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect \n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good \n", percent, "%")
	} else {
		fmt.Printf("%.1f%s bad \n", percent, "%")
	}

	fmt.Println("---------------------------------------")

	//control flow using switch case
	var anotherPointAgain = 6
	switch anotherPointAgain {
	case 8:
		fmt.Println("Perfect")
	case 7, 5, 6, 4, 3: //can using multiple conditions
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can better")
		}
	}

	switch {
	case anotherPointAgain == 10:
		fmt.Println("very very amazing")
	case (anotherPointAgain > 5) && (anotherPointAgain < 10):
		fmt.Println("Perfectss")
		fallthrough // ignore break in conditions
	case anotherPoint < 5:
		fmt.Println("awesomes k")
	default:
		{
			fmt.Println("not badks")
			fmt.Println("you can better")
		}
	}

	fmt.Println("------------------------//control-flow//---------------")
	fmt.Println("------------------------//looping//---------------")
	for i := 0; i < 5; i++ {
		fmt.Println("index : ", i)
	}
	fmt.Println("---------------------------------------")
	//for with conditions argument
	var i = 0
	for i < 5 {
		fmt.Println("index : ", i)
		i += 1 //i++
	}
	fmt.Println("---------------------------------------")
	var ii = 0
	for {
		fmt.Println("index : ", ii)

		ii++
		if ii == 5 {
			break
		}
	}
	fmt.Println("---------------------------------------")
	//for with break and continue
	for j := 0; j <= 10; j++ {
		if j%2 == 1 {
			continue
		}

		if j > 8 {
			break
		}

		fmt.Println("Angka", j)
	}
	fmt.Println("---------------------------------------")
	//for with nested loop
	for j := 0; j < 5; j++ {
		for k := j; k < 5; k++ {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------------")
	//for with nested loop give label
outerLoop:
	for j := 0; j < 5; j++ {
		for k := j; k < 5; k++ {
			if j == 3 {
				break outerLoop //back to outerLoop label
			}
			fmt.Print("matriks  [", j, "][", k, "]", "\n")
		}
	}

	fmt.Println("------------------------//looping//---------------")
	fmt.Println("------------------------//array//---------------")
	var someNames [4]string
	someNames[0] = "beddu"
	someNames[1] = "salam"
	someNames[2] = "abdul"
	someNames[3] = "petta"

	fmt.Println(someNames) //arrays

	fmt.Println(someNames[0], someNames[1], someNames[2], someNames[3])

	fmt.Println("---------------------------------------")
	var fruits = [4]string{"durian", "rambutan", "mangga", "apel"}
	fmt.Println("jumlah elemen array \t\t", len(fruits))
	fmt.Println("isi element \t\t\t", fruits)
	fmt.Println("---------------------------------------")
	var someNumber = [...]int{1, 2, 3, 4} //initialize without size explicity
	fmt.Println("jumlah elemen array \t\t", len(someNumber))
	fmt.Println("isi element \t\t\t", someNumber)
	fmt.Println("---------------------------------------")
	//var numbers1 = [2][3]int{[3]int{3, 2, 2}, [3]int{3, 4, 5}} -> redundant initilizer
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}} //we can using this for avoid redundant type
	//fmt.Println("numbers 1 : \t", numbers1)
	fmt.Println("numbers 2 : \t", numbers2)
	fmt.Println("---------------------------------------")
	for items := 0; items < len(fruits); items++ {
		fmt.Printf("index posisi : %d -> %s\n", items, fruits[items])
	}
	fmt.Println("---------------------------------------")
	//accessing all array element using range
	for indexs, fruit := range fruits { //change indexs to _ if just using fruit
		fmt.Printf("index posisi : %d -> %s\n", indexs, fruit)
	}

	//initialize array for allocate memory
	var someArrays = make([]int, 4)
	someArrays[0] = 1
	someArrays[1] = 2
	someArrays[2] = 3
	someArrays[3] = 4
	fmt.Println(someArrays)
	fmt.Println("------------------------//array//---------------")
	fmt.Println("------------------------//slice//---------------")
	var sliceFruits = []string{"apple", "grape", "banana", "melon"} //if sum of element not explicit declare, then the variable is slice
	fmt.Println(sliceFruits[0])
	var newSliceFruits = sliceFruits[0:2]
	fmt.Println(newSliceFruits)

	/*
		slice is reference type data
	*/

	fmt.Println("------------------------//slice//---------------")
	fmt.Println("------------------------//map//---------------")
	//var chicken map[string]int //declare map variable
	chicken := map[string]int{} // a map variable must be explicitly initialize because default value is nil

	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("Januari : ", chicken["januari"])
	fmt.Println("februari : ", chicken["februari"])
	fmt.Println("------------------------//map//---------------")

	var chicken1 = map[string]int{
		"january":  10,
		"february": 20,
	}

	fmt.Println("map :", chicken1)

	//another way declare map
	/*
		var chicken = map[string]int{}
		var chicken = make(map[string]int)
		var chicken = *new(map[string]int)
	*/

	//looping map
	var beef = map[string]int{
		"january":  10,
		"february": 30,
		"march":    20,
		"april":    40,
		"mei":      20,
	}

	for key, value := range beef {
		fmt.Println(key, "   \t: ", value)
	}

	//delete item map
	fmt.Println(len(beef))
	fmt.Print(beef)

	delete(beef, "mei")

	fmt.Println(len(beef))
	fmt.Println(beef)

	//detect if item is available

	var value, isAvailable = beef["january"]

	if isAvailable {
		fmt.Println(value)
	} else {
		fmt.Println("item isn't exist")
	}

}
