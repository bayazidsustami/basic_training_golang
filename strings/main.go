package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "jhon wick"
	var isExists = strings.Contains(str, "wick")
	fmt.Println(isExists)

	var isPrefix = strings.HasPrefix(str, "w")
	fmt.Println(isPrefix)

	var isSuffix = strings.HasSuffix(str, "ck")
	fmt.Println(isSuffix)

	var howMany = strings.Count(str, "k")
	fmt.Println(howMany)

	var index = strings.Index(str, "ic")
	fmt.Println(index)

	var find = "o"
	var replaceWith = "i"

	var newStr = strings.Replace(str, find, replaceWith, 1)
	fmt.Println(newStr)

	var strToRepeat = "no"
	fmt.Println(strings.Repeat(strToRepeat, 10))

	splitStr := strings.Split(str, " ")
	fmt.Println(splitStr)

	var data = []string{"banana", "papaya", "tomato"}
	var strFruits = strings.Join(data, "-")
	fmt.Println(strFruits)

	var toLower = strings.ToLower("aLaY")
	fmt.Println(toLower)

	var toUpper = strings.ToUpper("eat")
	fmt.Println(toUpper)
}
