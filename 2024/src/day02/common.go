package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read file and return a list of a list of integers
func readFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strNums := strings.Split(line, " ")
		var nums []int
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Println("Error converting string to int")
				return nil, err
			}
			nums = append(nums, num)
		}
		reports = append(reports, nums)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}
