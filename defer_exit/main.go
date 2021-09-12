package main

import "fmt"

func main() {

	orderSomeFood("pizza")
	orderSomeFood("burger")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih Silahkan Tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("pizza ditempat kami paling enak", "\n")
		return
	}
	fmt.Println("Pesanan anda : ", menu)
}
