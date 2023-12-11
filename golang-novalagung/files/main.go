package main

import (
	"fmt"
	"io"
	"os"
)

var path = "E:/project/test.txt"

func main() {
	createFile()
	writeFile()
	readFile()
	deleteFile()
}

func createFile() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("===> file berhasil dibuat", path)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func writeFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("file berhasil diisi")
}

func readFile() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}

	if isError(err) {
		return
	}

	fmt.Println("===> berhasil membaca file")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}
	fmt.Println("===> file berhasil dihapus")
}
