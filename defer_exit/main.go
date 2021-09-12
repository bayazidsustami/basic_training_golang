package main

import "fmt"

func main() {

	orderSomeFood("pizza")
	orderSomeFood("burger")

	fmt.Println("--------------------")
	deferAndIife()
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

//defer only for function scope
func deferAndIife() {
	number := 3
	if number == 3 {
		fmt.Println("halo 1")
		func() {
			defer fmt.Println("Halo 3")
		}()
	}
	fmt.Println("Halo 2")
}
