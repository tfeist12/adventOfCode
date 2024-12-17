package day08

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

// Calculate the distance between two positions
func calculateDistance(p1, p2 position) distance {
	return distance{p2.x - p1.x, p2.y - p1.y}
}

// Check if a position is within the bounds of the map
func isValidPosition(mapData [][]string, pos position) bool {
	return pos.x >= 0 && pos.x < len(mapData) && pos.y >= 0 && pos.y < len(mapData[0])
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
