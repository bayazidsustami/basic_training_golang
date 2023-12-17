package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for _, arg := range args {
		fmt.Println(arg)
	}

	hostName, err := os.Hostname()
	if err == nil {
		fmt.Println(hostName)
	} else {
		fmt.Println("err : ", err.Error())
	}
}
