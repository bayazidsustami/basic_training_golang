package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	createNewFile("sample.log", "this is sample log")
	appendToFile("sample.log", "appended")

	msg, err := readFile("sample.log")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(msg)
	}
}

func createNewFile(filename string, message string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(filename string) (string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		message += string(line) + "\n"
		if err == io.EOF {
			break
		}
	}
	return message, nil
}

func appendToFile(filename string, message string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}
