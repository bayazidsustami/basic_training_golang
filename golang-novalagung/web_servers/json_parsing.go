package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("User :", data.FullName)
	fmt.Println("age", data.Age)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("Name : ", data1["Name"])
	fmt.Println("age :", data1["Age"])

	decodeArrayJson()
	encodeToJsonString()
}

func decodeArrayJson() {
	var jsonString = `[
		{"Name": "Buys", "Age": 24},
		{"Name": "Bsa", "Age": 25}
		]`

	var data []User

	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User 1:", data[0].FullName)
	fmt.Println("User 2 :", data[1].FullName)
}

func encodeToJsonString() {
	var obj = []User{
		{"jhon", 21},
		{"wick", 22},
	}
	var jsonData, err = json.Marshal(obj)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
