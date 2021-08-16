package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

//closure scheme alias
type FilterCallback func(string) bool

func main() {
	var names = []string{"jhon", "wick"}
	printMessage("Hello", names)

	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number :", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number :", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number :", randomValue)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

	var diameter float64 = 15
	var area, circumReference = calculate(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumReference)

	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3} //using slice
	//var avg = calculateAverage(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var avg = calculateAverage(numbers...)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	//closure function / anonymous function
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	var anotherNumbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newAnotherNumbers = func(min int) []int {
		var r []int
		for _, e := range anotherNumbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	/*
		closure with manifest typing
		var closure (func (string, int, []string) int)
		closure = func (string, int, []string) int{
			.....
		}
	*/

	fmt.Println("Original Number :", anotherNumbers)
	fmt.Println("Filtered Number :", newAnotherNumbers)

	// closure function in return
	var howMany, getNumbers = findMax(numbers, 3)
	fmt.Println("found \t", howMany)
	fmt.Println("value \t", getNumbers())

	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)

	//print pointer
	printPointer()

	//pointer as params
	var originNumber = 4

	fmt.Println("before : ", originNumber)
	change(&originNumber, 10)
	fmt.Println("after : ", originNumber)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

//function with return value
func randomWithRange(min, max int) int {
	var value = rand.Int()%(min-max+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var result = m / n
	fmt.Printf("%d / %d == %d\n", m, n, result)
}

//function with multiple return
/*func calculate(d float64) (float64, float64) {
	//hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	var circumReference = math.Pi * d

	return area, circumReference
}*/

func calculate(d float64) (area float64, circumReference float64) { //predefine return value
	//hitung luas
	area = math.Pi * math.Pow(d/2, 2)
	circumReference = math.Pi * d

	return area, circumReference
}

//variadic argument
func calculateAverage(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

//function as parameter
func filter(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func printPointer() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("Number A (Value) : ", numberA)
	fmt.Println("NUmber A (Address) : ", &numberA)
	fmt.Println("Number B (Value) : ", *numberB)
	fmt.Println("NUmber B (Address) : ", numberB)

	fmt.Println("====================")

	numberA = 5

	fmt.Println("Number A (Value) : ", numberA)
	fmt.Println("NUmber A (Address) : ", &numberA)
	fmt.Println("Number B (Value) : ", *numberB)
	fmt.Println("NUmber B (Address) : ", numberB)
}

func change(original *int, value int) {
	*original = value
}
