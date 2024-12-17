package day06

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
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

// Print the map
func printMap(mapData [][]string) {
	for _, row := range mapData {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
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
