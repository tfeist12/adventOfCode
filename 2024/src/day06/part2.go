package day06

import (
	"fmt"
)

func copyMap(original [][]string) [][]string {
	mapDataCopy := make([][]string, len(original))
	for i := range original {
		mapDataCopy[i] = make([]string, len(original[i]))
		copy(mapDataCopy[i], original[i])
	}
	return mapDataCopy
}

// Check if adding an obstacle creates a loop
func createsLoop(mapData [][]string, obstaclePos position) bool {
	// Add the obstacle
	mapData[obstaclePos.y][obstaclePos.x] = obstacle

	// Get the initial position and direction of the guard
	guardPos, guard := getPosition(mapData, upGuard)
	guardDir := getDirection(guard)

	visitedStates := make(map[position]map[direction]bool)

	for {
		// Uncomment to view the guard moving
		// printMap(mapData)

		if !isOnMap(mapData, guardPos) {
			return false
		}
		if visitedStates[guardPos] == nil {
			visitedStates[guardPos] = make(map[direction]bool)
		}
		if visitedStates[guardPos][guardDir] {
			return true
		}
		visitedStates[guardPos][guardDir] = true

		guardPos, guardDir = moveGuard(mapData, guardPos, guardDir)
	}
}

func Part2(filename string) {
	mapData, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Save the original map
	mapDataCopy := copyMap(mapData)

	// Part 2: Count the number of positions where adding an obstace creates a loop
	loopObstacles := 0
	for i, row := range mapData {
		for j, char := range row {
			// Reset the map to the original state before next iteration
			mapData = copyMap(mapDataCopy)
			if char == empty {
				if createsLoop(mapData, position{j, i}) {
					loopObstacles++
				}
			}
		}
	}
	fmt.Printf(
		"There are '%d' positions where an obstacle can be added to create a loop\n",
		loopObstacles,
	)
}
