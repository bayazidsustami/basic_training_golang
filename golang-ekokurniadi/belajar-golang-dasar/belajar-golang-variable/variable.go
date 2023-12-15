package main

import "fmt"

func main() {
	var name string

	name = "bay"
	fmt.Println(name)

	name = "bayazid"
	fmt.Println(name)

	var number = 11
	fmt.Println(number)

	floating := 3.05 // bentuk lain dari var
	fmt.Println(floating)

	var (
		firstName = "bayazid"
		lastName  = "sustami"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	const versionName string = "1.0.0" //inisialisasi constant
	const versionCode = 1
	fmt.Println(versionName)
	fmt.Println(versionCode)

	const (
		buildType        string = "release"
		buildEnvironment string = "production"
	)
	fmt.Println(buildType)
	fmt.Println(buildEnvironment)
}
