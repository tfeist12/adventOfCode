package day06

import (
	"fmt"
	"strings"
)

// Count the number of visited positions on the map
func countVisited(mapData [][]string) int {
	count := 0
	for _, row := range mapData {
		for _, char := range row {
			if strings.ContainsAny(char, visited) {
				count++
			}
		}
	}
	return count
}

func Part1(filename string) {
	mapData, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Part 1: Count the number of visited positions
	guardPos, guard := getPosition(mapData, upGuard)
	guardDir := getDirection(guard)

	for {
		// Uncomment to view the guard moving
		// printMap(mapData)

		if !isOnMap(mapData, guardPos) {
			visitedPositions := countVisited(mapData)
			fmt.Printf("The guard visited '%d' positions\n", visitedPositions)
			break
		}
		guardPos, guardDir = moveGuard(mapData, guardPos, guardDir)
	}
}
