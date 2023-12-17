package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string `required:"true" max:"20"`
	Age  int    `required:"true" max:"5"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)

	fmt.Println("type name", valueType)

	for i := 0; i < valueType.NumField(); i++ {
		fieldValue := valueType.Field(i)
		fmt.Println("field name", fieldValue.Name)
		fmt.Println("field type", fieldValue.Type)
		fmt.Println("field tag required", fieldValue.Tag.Get("required"))
		fmt.Println("field tag max", fieldValue.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
		}
	}
	return result
}

func main() {
	readField(Sample{"bay"})
	readField(Person{"bayazid", 26})

	fmt.Println(isValid(Sample{"bay"}))
}
