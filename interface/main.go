package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var bangunDatar hitung
	bangunDatar = persegi{10.0}
	fmt.Println("=====persegi")
	fmt.Println("Luas		:", bangunDatar.luas())
	fmt.Println("Keliling	:", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("=====lingkaran")
	fmt.Println("Luas		:", bangunDatar.luas())
	fmt.Println("Keliling	:", bangunDatar.keliling())
	fmt.Println("jari-jari	:", bangunDatar.(lingkaran).jariJari())

	var bangunRuang hitung1d = &kubus{4}
	fmt.Println("=====lingkaran")
	fmt.Println("Luas		:", bangunRuang.luas())
	fmt.Println("keliling	:", bangunRuang.keliling())
	fmt.Println("Volume		:", bangunRuang.volume())

	fmt.Println("=====empty interface=====")
	var secret interface{}
	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	data := map[string]interface{}{
		"name":      "jhon ethan",
		"grade":     2,
		"breakfast": []string{"apple", "banana", "pineapple"},
	}

	fmt.Println(data["name"])
	fmt.Println(data["breakfast"])

	//casting empty variable interface
	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)
	secret = []string{"apple", "durian", "rambutan"}
	var fruits = strings.Join(secret.([]string), ",")
	fmt.Println(fruits, "is my favorite fruits")

	secret = &person{name: "charlie", age: 20}
	var name = secret.(*person).name
	fmt.Println(name, "was going to school")

	var persons = []map[string]interface{}{
		{"name": "wick", "age": 20},
		{"name": "jhon", "age": 19},
		{"name": "charlie", "age": 18},
	}

	for _, each := range persons {
		fmt.Println(each["name"], "age is : ", each["age"])
	}

	var fruites = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "apple", "papaya"},
		"orange",
	}

	for _, each := range fruites {
		fmt.Println(each)
	}
}

type person struct {
	name string
	age  int
}

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung1d interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}
