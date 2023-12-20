package belajar_golang_embed

import (
	_ "embed"
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
