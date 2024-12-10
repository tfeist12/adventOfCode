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

	empty    = "."
	visited  = "X"
	obstacle = "#"

	upGuard    = "^"
	rightGuard = ">"
	downGuard  = "v"
	leftGuard  = "<"
)

var (
	upDirection    = direction{dx: 0, dy: -1}
	rightDirection = direction{dx: 1, dy: 0}
	downDirection  = direction{dx: 0, dy: 1}
	leftDirection  = direction{dx: -1, dy: 0}
)

type position struct {
	x, y int
}

type direction struct {
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

// Get the direction of the guard
func getDirection(guard string) direction {
	switch guard {
	case upGuard:
		return upDirection
	case rightGuard:
		return rightDirection
	case downGuard:
		return downDirection
	case leftGuard:
		return leftDirection
	default:
		// Invalid direction
		return direction{dx: 0, dy: 0}
	}
}

// Get the guard character based on the direction
func getGuard(dir direction) string {
	switch dir {

	case upDirection:
		return upGuard
	case rightDirection:
		return rightGuard
	case downDirection:
		return downGuard
	case leftDirection:
		return leftGuard
	default:
		// Invalid direction
		return ""
	}
}

// Get the position of the guard on the map
func getPosition(mapData [][]string, guard string) (position, string) {
	for i, row := range mapData {
		for j, char := range row {
			if strings.ContainsAny(char, guard) {
				return position{j, i}, char
			}
		}
	}
	// Invalid position
	return position{-1, -1}, ""
}

// Rotate the guard to the right
func rotateRight(mapData [][]string, guardPos position, guardDir direction) direction {
	switch guardDir {
	case upDirection:
		mapData[guardPos.y][guardPos.x] = rightGuard
		return rightDirection
	case rightDirection:
		mapData[guardPos.y][guardPos.x] = downGuard
		return downDirection
	case downDirection:
		mapData[guardPos.y][guardPos.x] = leftGuard
		return leftDirection
	case leftDirection:
		mapData[guardPos.y][guardPos.x] = upGuard
		return upDirection
	default:
		// Invalid direction
		return direction{dx: 0, dy: 0}
	}
}

// Move the guard in the specified direction or rotate if they hit an obstacle
func moveGuard(mapData [][]string, pos position, dir direction) (position, direction) {
	newPos := position{pos.x + dir.dx, pos.y + dir.dy}

	if !isOnMap(mapData, newPos) {
		// Update current position to be visited
		mapData[pos.y][pos.x] = visited
		return newPos, dir
	} else if mapData[newPos.y][newPos.x] != obstacle {
		// Update the previous position to visited
		mapData[pos.y][pos.x] = visited
		// Update the new position to the guard
		mapData[newPos.y][newPos.x] = getGuard(dir)
		return newPos, dir
	} else {
		// Rotate the guard to the right
		newDir := rotateRight(mapData, pos, dir)
		return pos, newDir
	}
}

// Check if the position is on the map
func isOnMap(mapData [][]string, pos position) bool {
	if pos.y < 0 || pos.y >= len(mapData) {
		return false
	}
	if pos.x < 0 || pos.x >= len(mapData[0]) {
		return false
	}
	return true
}

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

func printMap(mapData [][]string) {
	for _, row := range mapData {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
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

func copyMap(original [][]string) [][]string {
	mapDataCopy := make([][]string, len(original))
	for i := range original {
		mapDataCopy[i] = make([]string, len(original[i]))
		copy(mapDataCopy[i], original[i])
	}
	return mapDataCopy
}

func main() {
	mapData, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Save the original map
	mapDataCopy := copyMap(mapData)

	// Part 1: Count the number of visited positions
	guardPos, guard := getPosition(mapData, upGuard)
	guardDir := getDirection(guard)

	for {
		// Uncomment to view the guard moving
		// printMap(mapData)

		if !isOnMap(mapData, guardPos) {
			visitedPositions := countVisited(mapData)
			fmt.Printf("The guard visited '%d' positions\n", visitedPositions)
			// Reset the map to the original state before part 2
			mapData = copyMap(mapDataCopy)
			break
		}
		guardPos, guardDir = moveGuard(mapData, guardPos, guardDir)
	}

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
