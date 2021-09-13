package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Println("what is 725/25 ?")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right")
	} else {
		fmt.Println("the answer is wrong")
	}
}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntimeout! no answer more than ", timeout, "seconds")
	os.Exit(0)
}
