package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
)

func main() {
	const FTP_ADDR = "0.0.0.0:2121"
	const FTP_USERNAME = "admin"
	const FTP_PASSWORD = "123456"

	connect, err := ftp.Dial(FTP_ADDR)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = connect.Login(FTP_USERNAME, FTP_PASSWORD)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("====== PATH /.")

	entries, err := connect.List(".")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, entry := range entries {
		fmt.Println(" ->", entry.Name, getStringEntryType(entry.Type))
	}
}

func getStringEntryType(t ftp.EntryType) string {
	switch t {
	case ftp.EntryTypeFile:
		return "(file)"
	case ftp.EntryTypeFolder:
		return "(folder)"
	case ftp.EntryTypeLink:
		return "(link)"
	default:
		return ""
	}
}
