package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var tempPath = filepath.Join(os.Getenv("TEMP"), "learn-pipeline-temp")

func main() {
	log.Println("start")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func proceed() {
	counterTotal := 0
	counterRename := 0

	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		counterTotal++

		//readFile
		buff, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		//sum
		sum := fmt.Sprintf("%x", md5.Sum(buff))

		destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.text", sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRename++
		return nil
	})

	if err != nil {
		log.Println("ERROR : ", err.Error())
	}

	log.Printf("%d/%d files renamed", counterRename, counterTotal)
}
