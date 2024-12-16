package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	exampleFilename = "example.txt"
	filename        = "input.txt"
)

type position struct {
	x, y int
}

type distance struct {
	dx, dy int
}

// Read the map file and return a 2D list of strings for positions
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

// Calculate the distance between two positions
func calculateDistance(p1, p2 position) distance {
	return distance{p2.x - p1.x, p2.y - p1.y}
}

// Check if a position is within the bounds of the map
func isValidPosition(mapData [][]string, pos position) bool {
	return pos.x >= 0 && pos.x < len(mapData) && pos.y >= 0 && pos.y < len(mapData[0])
}

// Create an empty map of the same size as the input map
func createEmptyMap(mapData [][]string) [][]string {
	emptyMap := make([][]string, len(mapData))
	for i := range mapData {
		emptyMap[i] = make([]string, len(mapData[i]))
		for j := range mapData[i] {
			emptyMap[i][j] = "."
		}
	}
	return emptyMap
}

// Return a map of antinode positions
func getAntiNodeMap(mapData [][]string) [][]string {
	antiNodeMap := createEmptyMap(mapData)

	// Iterate over the map
	for i, row := range mapData {
		for j, char := range row {
			// Only iterate if character is not an empty space
			if char != "." {
				// Find distance to the same character
				for k, row2 := range mapData {
					for l, col2 := range row2 {
						if col2 == char && (i != k || j != l) {
							dist := calculateDistance(position{i, j}, position{k, l})
							// Check for an open position that is the same distance away from the second position
							openPos := position{k + dist.dx, l + dist.dy}
							if isValidPosition(mapData, openPos) &&
								antiNodeMap[openPos.x][openPos.y] != "#" {
								antiNodeMap[openPos.x][openPos.y] = "#"
							}
						}
					}
				}
			}
		}
	}
	return antiNodeMap
}

// Return a map of antinode positions
// This version uses T frequency signals for the model
func getAntiNodeMapT(mapData [][]string) [][]string {
	antiNodeMap := createEmptyMap(mapData)

	// Iterate over the map
	for i, row := range mapData {
		for j, col := range row {
			if col != "." && col != "#" {
				char := col
				for k, row2 := range mapData {
					for l, col2 := range row2 {
						if col2 == char && (i != k || j != l) {
							antiNodeMap[i][j] = "#"
							dist := calculateDistance(position{i, j}, position{k, l})
							// Check for open positions beyond the second antenna
							multiplier := 1
							for {
								openPos := position{k + dist.dx*multiplier, l + dist.dy*multiplier}
								if !isValidPosition(mapData, openPos) {
									break
								}
								if antiNodeMap[openPos.x][openPos.y] == "." {
									antiNodeMap[openPos.x][openPos.y] = "#"
								}
								multiplier++
							}
						}
					}
				}
			}
		}
	}
	return antiNodeMap
}

// Count all instances of '#' in the map
func countAntiNodePositions(mapData [][]string) int {
	count := 0
	for _, row := range mapData {
		for _, char := range row {
			if char == "#" {
				count++
			}
		}
	}
	return count
}

// Print the map
func printMap(mapData [][]string) {
	for _, row := range mapData {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
}

func main() {
	mapData, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Part 1
	antiNodeMap := getAntiNodeMap(mapData)
	count := countAntiNodePositions(antiNodeMap)
	fmt.Printf("Number of antinode positions: %d\n", count)

	// Part 2
	antiNodeMap = getAntiNodeMapT(mapData)
	count = countAntiNodePositions(antiNodeMap)
	fmt.Printf("Number of antinode positions using T frequency: %d\n", count)
}
