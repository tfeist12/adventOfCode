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

func countAntinodePositions(mapData [][]string) (int, [][]string) {
	antiNodeMap := make([][]string, len(mapData))
	for i := range mapData {
		antiNodeMap[i] = make([]string, len(mapData[i]))
		for j := range mapData[i] {
			antiNodeMap[i][j] = "."
		}
	}

	count := 0
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
								count++
							}
						}
					}
				}
			}
		}
	}
	return count, antiNodeMap
}

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

	count, antiNodeMap := countAntinodePositions(mapData)
	fmt.Printf("Number of antinode positions: %d\n", count)
	printMap(antiNodeMap)
}
