package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "212"
	num, err := strconv.Atoi(str)

	if err == nil {
		fmt.Println(num)
	}

	var numbers = 202
	strr := strconv.Itoa(numbers)

	fmt.Println(strr)

	parseInt()
	parseFloat()
	parseBool()
	convertWithCasting()
	stringToByte()
	typeAssertions()
}

func parseInt() {
	var str = "1010"
	var num, err = strconv.ParseInt(str, 2, 8)
	if err == nil {
		fmt.Println(num)
	}

	var number = int64(24)
	var strings = strconv.FormatInt(number, 8)
	fmt.Println(strings)
}

func parseFloat() {
	var str = "24.12"
	num, err := strconv.ParseFloat(str, 32)
	if err == nil {
		fmt.Println(num)
	}

	var numbers = float64(24.12)
	var strings = strconv.FormatFloat(numbers, 'f', 6, 64)
	fmt.Println(strings)
}

func parseBool() {
	var str = "true"
	var bul, err = strconv.ParseBool(str)

	if err == nil {
		fmt.Println(bul)
	}

	var buls = false
	var strings = strconv.FormatBool(buls)
	fmt.Println(strings)
}

func convertWithCasting() {
	var a float64 = float64(24)
	fmt.Println(a)

	var b int32 = int32(24.00)
	fmt.Println(b)
}

func stringToByte() {
	var str = "halo"
	var b = []byte(str)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3]) //output ASCII Code

	var bytes = []byte{104, 97, 108, 111}
	var strings = string(bytes)

	fmt.Printf("%s \n", strings)

	var c int64 = int64('o')
	fmt.Println(c)

	var d string = string(111)
	fmt.Println(d)
}

func typeAssertions() {
	var data = map[string]interface{}{
		"nama":    "jhon wick",
		"grade":   2,
		"height":  156.2,
		"isMale":  true,
		"hobbies": []string{"eating, sleeping"},
	}

	fmt.Println(data["nama"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	//if you don't know what actual type of interface
	for _, value := range data {
		switch value.(type) {
		case string:
			fmt.Println(value.(string))
		case int:
			fmt.Println(value.(int))
		case float64:
			fmt.Println(value.(float64))
		case bool:
			fmt.Println(value.(bool))
		case []string:
			fmt.Println(value.([]string))
		default:
			fmt.Println(value.(int))
		}
	}
}
