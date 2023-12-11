package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	var text = "this is secret text"
	var sha = sha1.New()

	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)

	var hashed1, _ = doHashUsingSalt(text)
	fmt.Printf("hashed1 : %s\n\n", hashed1)

	var hashed2, _ = doHashUsingSalt(text)
	fmt.Printf("hashed2 : %s\n\n", hashed2)
}

func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text : '%s' \nsalt: %s", text, salt)
	fmt.Println(saltedText)

	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)
	return fmt.Sprintf("%x", encrypted), salt
}
