package day08

import (
	"fmt"
)

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

func Part1(filename string) {
	mapData, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	antiNodeMap := getAntiNodeMap(mapData)
	count := countAntiNodePositions(antiNodeMap)
	fmt.Printf("Number of antinode positions: %d\n", count)
}
