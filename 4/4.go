package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	filename = "input.txt"
	xmas     = "XMAS"
)

var directions = [][]int{
	{0, 1},   // horizontal right
	{0, -1},  // horizontal left
	{1, 0},   // vertical down
	{-1, 0},  // vertical up
	{1, 1},   // diagonal down right
	{1, -1},  // diagonal down left
	{-1, 1},  // diagonal up right
	{-1, -1}, // diagonal up left
}

var patterns = [][][]string{
	{
		{"M", ".", "S"},
		{".", "A", "."},
		{"M", ".", "S"},
	},
	{
		{"S", ".", "M"},
		{".", "A", "."},
		{"S", ".", "M"},
	},
	{
		{"M", ".", "M"},
		{".", "A", "."},
		{"S", ".", "S"},
	},
	{
		{"S", ".", "S"},
		{".", "A", "."},
		{"M", ".", "M"},
	},
}

// Read file in as a 2D list
func readFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var elements []string
		for _, char := range line {
			elements = append(elements, string(char))
		}
		result = append(result, elements)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func checkXmasDirection(grid [][]string, y, x, dy, dx int) bool {
	rows := len(grid)
	cols := len(grid[0])

	for k := 0; k < len(xmas); k++ {
		// Next positions
		ny := y + k*dy
		nx := x + k*dx
		// Return false if new position is out of bounds
		if ny < 0 || ny >= rows || nx < 0 || nx >= cols {
			return false
		}
		// Return false if character at new position doesn't match the current expected character
		if grid[ny][nx] != string(xmas[k]) {
			return false
		}
	}
	return true
}

func countXmas(grid [][]string) int {
	count := 0
	// Iterate through rows
	for i := 0; i < len(grid); i++ {
		// Iterate through columns
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range directions {
				if checkXmasDirection(grid, i, j, dir[0], dir[1]) {
					count += 1
				}
			}
		}
	}
	return count
}

func checkXmasPattern(grid [][]string, y, x int, pattern [][]string) bool {
	rows := len(grid)
	cols := len(grid[0])
	patternRows := len(pattern)
	patternCols := len(pattern[0])

	// Iterate through the pattern
	for i := 0; i < patternRows; i++ {
		for j := 0; j < patternCols; j++ {
			// Calculate new indices in the grid
			ny := y + i
			nx := x + j
			// Return false if new position is out of bounds
			if ny < 0 || ny >= rows || nx < 0 || nx >= cols {
				return false
			}
			if pattern[i][j] != "." && pattern[i][j] != grid[ny][nx] {
				return false
			}
		}
	}
	return true
}

func countXmasPattern(grid [][]string) int {
	count := 0
	// Iterate through rows
	for i := 0; i < len(grid); i++ {
		// Iterate through columns
		for j := 0; j < len(grid[i]); j++ {
			// Iterate through valid patterns
			for _, pattern := range patterns {
				if checkXmasPattern(grid, i, j, pattern) {
					count += 1
				}
			}
		}
	}
	return count
}

func main() {
	puzzleGrid, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Part 1
	occurances := countXmas(puzzleGrid)
	fmt.Printf("Number of 'XMAS' instances in the word search: %d\n", occurances)

	// Part 2
	patterns := countXmasPattern(puzzleGrid)
	fmt.Printf("Number of 'X-MAS' patterns in the word search: %d\n", patterns)
}
