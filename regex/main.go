package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)
	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	var str = regex.FindString(text)
	fmt.Println(str)

	var index = regex.FindStringIndex(text)
	fmt.Println(index)

	var strIdx = text[0:6]
	fmt.Println(strIdx)

	var strToRepl = regex.ReplaceAllStringFunc(text, func(s string) string {
		if s == "burger" {
			return "potato"
		}
		return s
	})

	fmt.Println(strToRepl)

	var textRepl = regex.ReplaceAllString(text, "potato")
	fmt.Println(textRepl)

}
