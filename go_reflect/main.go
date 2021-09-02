package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("tipe variable :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable : ", reflectValue.Int())
	}

	fmt.Println("nilai varible : ", reflectValue.Interface())

	var s1 = &student{Name: "wick", Age: 20}
	s1.getPropertyInfo()
}

type student struct {
	Name string
	Age  int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("name	:", reflectType.Field(i).Name)
		fmt.Println("Tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai	:", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}
