package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test_case/word_search_test_case.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var T int
	if scanner.Scan() {
		fmt.Sscan(scanner.Text(), &T)
	}
	for t := 0; t < T && scanner.Scan(); t++ {
		var N, M int
		fmt.Sscan(scanner.Text(), &N)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &M)

		grid := make([]string, N)
		for i := 0; i < N && scanner.Scan(); i++ {
			grid[i] = scanner.Text()
		}
		scanner.Scan()
		word := scanner.Text()
		fmt.Printf("Case %d:%d\n", t+1, countWordOccurrences(grid, word))
	}
}

func countWordOccurrences(grid []string, word string) int {
	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			count += countFromPosition(grid, word, r, c)
		}
	}
	return count
}

func countFromPosition(grid []string, word string, row, col int) int {
	count := 0
	// Check horizontal right
	if col+len(word) <= len(grid[0]) && grid[row][col:col+len(word)] == word {
		count++
	}
	// Check horizontal left
	if col-len(word)+1 >= 0 {
		reversedWord := reverseString(word)
		if grid[row][col-len(word)+1:col+1] == reversedWord {
			count++
		}
	}
	// Check vertical down
	if row+len(word) <= len(grid) {
		match := true
		for i := 0; i < len(word); i++ {
			if grid[row+i][col] != word[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	// Check vertical up
	if row-len(word)+1 >= 0 {
		reversedWord := reverseString(word)
		match := true
		for i := 0; i < len(word); i++ {
			if grid[row-i][col] != reversedWord[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	// Check diagonal top-left to bottom-right
	if row+len(word) <= len(grid) && col+len(word) <= len(grid[0]) {
		match := true
		for i := 0; i < len(word); i++ {
			if grid[row+i][col+i] != word[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	// Check diagonal top-right to bottom-left
	if row-len(word)+1 >= 0 && col+len(word) <= len(grid[0]) {
		reversedWord := reverseString(word)
		match := true
		for i := 0; i < len(word); i++ {
			if grid[row-i][col+i] != reversedWord[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	return count
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
