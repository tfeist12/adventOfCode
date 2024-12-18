package day09

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	emptyBlock = -1
)

// Read full file into a string
func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Remove all characters from a string that aren't digits
func removeNonDigits(input string) string {
	var result strings.Builder
	for _, char := range input {
		if unicode.IsDigit(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

// Convert the raw data string into a block map list
func buildBlockMap(data string) ([]int, error) {
	data = removeNonDigits(data)

	var blockMap []int

	// Iterate through the data string
	for i := 0; i < len(data); i++ {
		// Each digit is the number of spaces in the block map
		spaces, err := strconv.Atoi(string(data[i]))
		if err != nil {
			fmt.Printf("Error converting '%s' to an int\n", string(data[i]))
			return nil, err
		}
		// If the index is even, use the file ID for the number of spaces
		// If the index is odd, use the empty char for the number of spaces
		if i%2 == 0 {
			fileID := i / 2
			for j := 0; j < spaces; j++ {
				blockMap = append(blockMap, fileID)
			}
		} else {
			for j := 0; j < spaces; j++ {
				blockMap = append(blockMap, emptyBlock)
			}
		}
	}
	return blockMap, nil
}

// Calculate the checksum of the compacted block map
func calcCompactedChecksum(compactedBlockMap []int) int {
	checkSum := 0
	for i := 0; i < len(compactedBlockMap); i++ {
		if compactedBlockMap[i] != emptyBlock {
			checkSum += i * compactedBlockMap[i]
		}
	}

	return checkSum
}
