package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("bayazid"))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	csvReader()
	csvWriter()
}

func csvReader() {
	csvString := "bay, yazid, sustami\n" +
		"budi, pratama, utama\n" +
		"joko, morro, diah"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}

func csvWriter() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"bay", "yazid", "sustami"})
	_ = writer.Write([]string{"budi", "pratama", "utama"})
	_ = writer.Write([]string{"joko", "morro", "diah"})
	writer.Flush()
}
