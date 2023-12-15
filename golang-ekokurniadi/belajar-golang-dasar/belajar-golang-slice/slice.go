package main

import "fmt"

func main() {
	names := [...]string{
		"Eko",
		"Kurniawan",
		"Khanedy",
		"Joko",
		"Budi",
		"Nugraha",
	}

	slice1 := names[4:6]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1)) //capacity

	slice2 := names[:4]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)

	daysSlice := days[5:]
	daysSlice[0] = "Sabtu baru"
	daysSlice[1] = "Minggu baru"
	fmt.Println(days) //data yang di ubah merujuk ke data array

	daysSlice1 := append(daysSlice, "Liburr")
	daysSlice1[0] = "ups"
	fmt.Println(daysSlice1)
	fmt.Println(days) // data days tidak berubah setelah append karena append membuat array baru

	newSlice := make([]string, 2, 5)
	newSlice[0] = "bay"
	newSlice[1] = "baya"
	// newSlice[2] = "test" --> error karena size nya sudah di set 2, harus menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "test")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Joko"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	originSlice := days[:]
	destinationSlice := make([]string, len(originSlice), cap(originSlice))

	copy(destinationSlice, originSlice)
	fmt.Println(originSlice)
	fmt.Println(destinationSlice)

	exampleArray := [...]int{1, 2, 3, 4, 5} // array harus didefinisikan sizenya
	exampleSlice := []int{1, 2, 3, 4, 5}    // slice tidak harus didefinisikan

	fmt.Println(exampleArray)
	fmt.Println(exampleSlice)
}
