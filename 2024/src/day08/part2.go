package day08

import (
	"fmt"
)

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

func Part2(filename string) {
	mapData, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	antiNodeMap := getAntiNodeMapT(mapData)
	count := countAntiNodePositions(antiNodeMap)
	fmt.Printf("Number of antinode positions using T frequency: %d\n", count)
}
