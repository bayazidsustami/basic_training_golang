package belajar_golang_embed

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}

//go:embed lorem.png
var logo []byte

func TestEmbedBinary(t *testing.T) {
	err := os.WriteFile("new_logo.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedMultipleFiles(t *testing.T) {
	fileA, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(fileA))

	fileB, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(fileB))

	fileC, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(fileC))
}
