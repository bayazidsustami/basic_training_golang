package library

import "fmt"

//access level public or exported using upper case
func SayHello(name string) {
	fmt.Println("hello")
	introduce(name)
}

//access level private unexported
func introduce(name string) {
	fmt.Println("my name :", name)
}

type Student struct {
	Name  string
	Grade int
}

func init() {
	fmt.Println("-----> library/library.go imported")
}
