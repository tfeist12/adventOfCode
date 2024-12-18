package day09

import (
	"fmt"
)

// Compact the block map by swapping the right most digit with the left most emptyBlock
func compactBlockMapIndiv(blockMap []int) []int {
	for {
		firstEmptyIndex := -1
		for i, v := range blockMap {
			if v == emptyBlock {
				firstEmptyIndex = i
				break
			}
		}
		// No empty blocks, exit the loop
		if firstEmptyIndex == -1 {
		}

		lastDigitIndex := -1
		for i := len(blockMap) - 1; i >= 0; i-- {
			if blockMap[i] != emptyBlock {
				lastDigitIndex = i
				break
			}
		}
		// No digits found or no more swaps needed
		if lastDigitIndex == -1 || lastDigitIndex < firstEmptyIndex {
			break
		}

		// Swap the final digit and the first emptyBlock
		blockMap[firstEmptyIndex], blockMap[lastDigitIndex] = blockMap[lastDigitIndex], blockMap[firstEmptyIndex]
	}
	return blockMap
}

func Part1(filename string) {
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

	compactedBlockMap := compactBlockMapIndiv(blockMap)

	checksum := calcCompactedChecksum(compactedBlockMap)
	fmt.Printf("Compacted checksum: %d\n", checksum)
}
