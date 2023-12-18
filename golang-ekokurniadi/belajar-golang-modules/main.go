package main

import (
	"fmt"

	go_say_hello "github.com/ProgrammerZamanNow/go-say-hello/v2"
)

func main() {
	result := go_say_hello.SayHello("bays")
	fmt.Println(result)
}
