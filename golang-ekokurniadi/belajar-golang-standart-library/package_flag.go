package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "put your database host")
	port := flag.Int("port", 0, "put your host port")
	username := flag.String("username", "root", "put your database username")
	password := flag.String("password", "root", "put your database password")

	flag.Parse()

	fmt.Println(*host, *port, *username, *password)
}
