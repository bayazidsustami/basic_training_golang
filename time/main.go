package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time %v\n", time1)

	var time2 = time.Date(2021, 11, 23, 11, 0, 0, 0, time.UTC)
	fmt.Printf("time %v\n", time2)

	fmt.Println("Year :", time1.Year(), "month:", time1.Month())
}
