package day04

import (
	"fmt"
)

const (
	xmas = "XMAS"
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

func Part1(filename string) {
	puzzleGrid, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	occurances := countXmas(puzzleGrid)
	fmt.Printf("Number of 'XMAS' instances in the word search: %d\n", occurances)
}
