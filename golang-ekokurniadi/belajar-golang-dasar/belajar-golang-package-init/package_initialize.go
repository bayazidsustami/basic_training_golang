package main

import (
	"belajar-golang-package-init/database"
	_ "belajar-golang-package-init/internal" // auto run init function
	"fmt"
)

func main() {
	databaseName := database.GetDatabase()
	fmt.Println(databaseName)
}
