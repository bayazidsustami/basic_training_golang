package main

import (
	"fmt"
	"math"
)

func main() {
	var bangunDatar hitung
	bangunDatar = persegi{10.0}
	fmt.Println("=====persegi")
	fmt.Println("Luas		:", bangunDatar.luas())
	fmt.Println("Keliling	:", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("=====lingkarang")
	fmt.Println("Luas		:", bangunDatar.luas())
	fmt.Println("Keliling	:", bangunDatar.keliling())
	fmt.Println("jari-jari	:", bangunDatar.(lingkaran).jariJari())
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
