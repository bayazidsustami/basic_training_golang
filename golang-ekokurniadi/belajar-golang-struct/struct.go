package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer
	customer.Name = "bay"
	customer.Address = "mks"
	customer.Age = 26

	fmt.Println(customer)
	fmt.Println(customer.Name)

	customer1 := Customer{
		Name:    "customer1",
		Address: "address1",
		Age:     0,
	}
	fmt.Println(customer1)

	customer1.sayHello()

	customer2 := Customer{"Customer2", "adress2", 0}
	fmt.Println(customer2)

	customer2.sayHello()
}

// struct method
func (customer Customer) sayHello() {
	fmt.Println("hello my nameIs", customer.Name)
}
