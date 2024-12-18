package day09

import (
	"fmt"
	"sort"
)

// Get a map of file IDs to their lengths
func getAllFiles(blockMap []int) map[int]int {
	fileLengths := make(map[int]int)
	n := len(blockMap)
	for i := 0; i < n; i++ {
		if blockMap[i] != -1 {
			fileID := blockMap[i]
			fileLength := 0
			for i < n && blockMap[i] == fileID {
				fileLength++
				i++
			}
			fileLengths[fileID] = fileLength
			i--
		}
	}
	return fileLengths
}

// Find the leftmost free space before a file that would fit it
func findLeftmostFreeSpace(blockMap []int, length int, end int) int {
	for i := 0; i <= end-length; i++ {
		free := true
		for j := 0; j < length; j++ {
			if blockMap[i+j] != -1 {
				free = false
				break
			}
		}
		if free {
			return i
		}
	}
	return -1
}

// Swap two ranges of equal length in the block map
func swapRanges(blockMap []int, start1, end1, start2, end2 int) {
	if end1-start1 != end2-start2 {
		fmt.Println("Ranges are not of equal length, cannot swap")
		return
	}

	for i := 0; i <= end1-start1; i++ {
		blockMap[start1+i] = blockMap[start2+i]
		blockMap[start2+i] = blockMap[start1+i]
	}
}

// Get the start index of a file in the block map
func getFileStartIndex(blockMap []int, fileID int) int {
	for i, id := range blockMap {
		if id == fileID {
			return i
		}
	}
	// If the file ID is not found, return -1
	return -1
}

// Move full files to the left most free space that will fit them
func compactFiles(blockMap []int) {
	fileLengths := getAllFiles(blockMap)
	fileIDs := make([]int, 0, len(fileLengths))
	for fileID := range fileLengths {
		fileIDs = append(fileIDs, fileID)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(fileIDs)))

	// Iterate over the file IDs in descending order
	for _, fileID := range fileIDs {
		fileLength := fileLengths[fileID]
		fileStart := getFileStartIndex(blockMap, fileID)

		// Determine the leftmost free space before the file
		// If the file fits in the free space, make a swap
		leftmostFreeSpace := findLeftmostFreeSpace(blockMap, fileLength, fileStart)
		if leftmostFreeSpace != -1 {
			swapRanges(
				blockMap,
				fileStart,
				fileStart+fileLength-1,
				leftmostFreeSpace,
				leftmostFreeSpace+fileLength-1,
			)
		}
	}
}

func printBlockMap(blockMap []int) {
	for _, block := range blockMap {
		if block == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(block)
		}
	}
	fmt.Println()
}

func Part2(filename string) {
	data, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	blockMap, err := buildBlockMap(data)
	if err != nil {
		fmt.Println("Error building block map")
		return
	}

	compactFiles(blockMap)
	checksum := calcCompactedChecksum(blockMap)
	fmt.Printf("Compacted checksum: %d\n", checksum)
}
