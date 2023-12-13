package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	for counter1 := 1; counter1 <= 5; counter1++ {
		fmt.Println("perulangan ke", counter1)
	}

	names := []string{"baya", "yazid", "sustami"}
	for i := 0; i < len(names); i++ {
		fmt.Println("name :", names[i])
	}

	for index, name := range names {
		fmt.Println("name :", name, "-", "index", index)
	}

	// break
	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("perulangan ke ", i)
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan ke ", i)
	}
}
