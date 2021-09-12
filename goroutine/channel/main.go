package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayToHello = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayToHello("jhon wick")
	go sayToHello("ethan hunt")
	go sayToHello("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	var msgForPrint = make(chan string)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			msgForPrint <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		printMessage(msgForPrint)
	}
}

//channel as parameter
func printMessage(what chan string) {
	fmt.Println(<-what)
}
