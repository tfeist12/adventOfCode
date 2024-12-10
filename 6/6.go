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

// Move the guard up or rotate if they hit an obstacle
func moveUp(mapData [][]string, pos position) (position, direction) {
	newPos := position{pos.x, pos.y - 1}

	if !isOnMap(mapData, newPos) {
		// Update current position to be visited
		mapData[pos.y][pos.x] = visited
		return newPos, upDirection
	} else if mapData[newPos.y][newPos.x] != obstacle {
		// Update the previous position to visited
		mapData[pos.y][pos.x] = visited
		// Update the new position to the guard
		mapData[newPos.y][newPos.x] = upGuard
		return newPos, upDirection
	} else {
		// Rotate the guard to the right
		newDir := rotateRight(mapData, pos, upDirection)
		return pos, newDir
	}
}

// Move the guard right or rotate if they hit an obstacle
func moveRight(mapData [][]string, pos position) (position, direction) {
	newPos := position{pos.x + 1, pos.y}

	if !isOnMap(mapData, newPos) {
		// Update current position to be visited
		mapData[pos.y][pos.x] = visited
		return newPos, rightDirection
	} else if mapData[newPos.y][newPos.x] != obstacle {
		// Update the previous position to visited
		mapData[pos.y][pos.x] = visited
		// Update the new position to the guard
		mapData[newPos.y][newPos.x] = rightGuard
		return newPos, rightDirection
	} else {
		// Rotate the guard to the right
		newDir := rotateRight(mapData, pos, rightDirection)
		return pos, newDir
	}
}

// Move the guard down or rotate if they hit an obstacle
func moveDown(mapData [][]string, pos position) (position, direction) {
	newPos := position{pos.x, pos.y + 1}

	if !isOnMap(mapData, newPos) {
		// Update current position to be visited
		mapData[pos.y][pos.x] = visited
		return newPos, downDirection
	} else if mapData[newPos.y][newPos.x] != obstacle {
		// Update the previous position to visited
		mapData[pos.y][pos.x] = visited
		// Update the new position to the guard
		mapData[newPos.y][newPos.x] = downGuard
		return newPos, downDirection
	} else {
		// Rotate the guard to the right
		newDir := rotateRight(mapData, pos, downDirection)
		return pos, newDir
	}
}

// Move the guard left or rotate if they hit an obstacle
func moveLeft(mapData [][]string, pos position) (position, direction) {
	newPos := position{pos.x - 1, pos.y}

	if !isOnMap(mapData, newPos) {
		// Update current position to be visited
		mapData[pos.y][pos.x] = visited
		return newPos, leftDirection
	} else if mapData[newPos.y][newPos.x] != obstacle {
		// Update the previous position to visited
		mapData[pos.y][pos.x] = visited
		// Update the new position to the guard
		mapData[newPos.y][newPos.x] = leftGuard
		return newPos, leftDirection
	} else {
		// Rotate the guard to the right
		newDir := rotateRight(mapData, pos, leftDirection)
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
}

func main() {
	mapData, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Get the inital position and direction of the guard
	guardPos, guard := getPosition(mapData, upGuard)
	guardDir := getDirection(guard)

	for {
		// Uncomment to view the guard moving
		// printMap(mapData)
		// fmt.Println()

		if !isOnMap(mapData, guardPos) {
			// fmt.Println("Left the map!")
			visitedPositions := countVisited(mapData)
			fmt.Printf("The guard visited '%d' positions\n", visitedPositions)
			return
		}
		if guardDir == upDirection {
			// fmt.Println("Moving up")
			guardPos, guardDir = moveUp(mapData, guardPos)
			continue
		}
		if guardDir == rightDirection {
			// fmt.Println("Moving right")
			guardPos, guardDir = moveRight(mapData, guardPos)
			continue
		}
		if guardDir == downDirection {
			// fmt.Println("Moving down")
			guardPos, guardDir = moveDown(mapData, guardPos)
			continue
		}
		if guardDir == leftDirection {
			// fmt.Println("Moving left")
			guardPos, guardDir = moveLeft(mapData, guardPos)
			continue
		}
	}
}
