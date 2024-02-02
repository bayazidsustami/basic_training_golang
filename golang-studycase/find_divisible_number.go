package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("test_case/divisible_number_test_case.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numIterations, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error reading number of iterations:", err)
		return
	}

	for i := 0; i < numIterations; i++ {
		var values []int
		for j := 0; j < 3; j++ {
			scanner.Scan()
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error reading value:", err)
				return
			}
			values = append(values, val)
		}

		fmt.Println("Case ", i+1, ":", countDivisibleNumbers(values[0], values[1], values[2]))
	}
}

func countDivisibleNumbers(a, b, k int) int {
	count := 0
	for num := a; num <= b; num++ {
		if num%k == 0 {
			count++
		}
	}
	return count
}
