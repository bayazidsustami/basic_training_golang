package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2023, time.August, 23, 12, 30, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	fmt.Println(utc.Year())
	fmt.Println(utc.Month())
	fmt.Println(utc.Day())

	formatter := "2006-01-02 15:04:05"
	parse, _ := time.Parse(formatter, "2020-12-10 12:30:10")
	fmt.Println(parse.Local())

	var duration1 time.Duration = time.Second * 100 //100 second
	duration2 := time.Millisecond * 10
	duration3 := duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration %d\n", duration3)
}
