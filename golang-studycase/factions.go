package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	x, y int
}

func isInsideGrid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func dfs(grid [][]byte, visited [][]bool, x, y int, letter byte, rows, cols int, factionMap map[byte]int) {
	if !isInsideGrid(x, y, rows, cols) || visited[x][y] || grid[x][y] == '#' || grid[x][y] != letter {
		return
	}

	visited[x][y] = true
	factionMap[letter]++

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		dfs(grid, visited, nx, ny, letter, rows, cols, factionMap)
	}
}

func findRegions(grid [][]byte, rows, cols int) (map[byte]int, int) {
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	factionMap := make(map[byte]int)
	contested := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] && grid[i][j] != '#' {
				letter := grid[i][j]
				dfs(grid, visited, i, j, letter, rows, cols, factionMap)
				if factionMap[letter] > 1 {
					contested++
				}
			}
		}
	}

	return factionMap, contested
}

func main() {
	file, err := os.Open("test_case/factions_test_case.in")
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

	for t := 1; t <= numIterations; t++ {
		scanner.Scan()
		var N, M int
		fmt.Sscanf(scanner.Text(), "%d", &N)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &M)

		grid := make([][]byte, N)
		for i := range grid {
			scanner.Scan()
			grid[i] = []byte(scanner.Text())
		}

		factionMap, contested := findRegions(grid, N, M)

		appendToFile("test_case/factions_test_case.out", fmt.Sprintf("Case %d:\n", t))

		factions := make([]byte, 0, len(factionMap))
		for faction := range factionMap {
			factions = append(factions, faction)
		}
		sort.Slice(factions, func(i, j int) bool { return factions[i] < factions[j] })

		for _, faction := range factions {
			fmt.Printf("%c %d\n", faction, factionMap[faction])
			appendToFile("test_case/factions_test_case.out", fmt.Sprintf("%c %d\n", faction, factionMap[faction]))
		}

		appendToFile("test_case/factions_test_case.out", fmt.Sprintf("contested %d\n", contested))
	}
}

func appendToFile(filename string, message string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}
