// Copyright © 2024 Tyler Feist tfeist612@gmail.com

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"aoc/src/day01"
	"aoc/src/day02"
	"aoc/src/day03"
	"aoc/src/day04"
	"aoc/src/day05"
	"aoc/src/day06"
	"aoc/src/day07"
	"aoc/src/day08"
	"aoc/src/day09"
)

var (
	day, part int
	fileType  string
)

// runnerCmd represents the runner command
var runnerCmd = &cobra.Command{
	Use:   "runner",
	Short: "Run a specific Advent of Code solution",
	Long:  "Run a specific Advent of Code solution for a given day, part, and input file.",
	Run: func(cmd *cobra.Command, args []string) {
		if day < 1 || day > 25 {
			log.Fatal("Invalid day. Use --day=1-25.")
		}
		if part != 1 && part != 2 {
			log.Fatal("Invalid part. Use --part=1 or --part=2.")
		}
		if fileType != "example" && fileType != "input" {
			log.Fatal("Invalid file. Use --file=example or --file=input.")
		}
		file := fmt.Sprintf("input/day%02d/%s.txt", day, fileType)
		if !fileExists(file) {
			log.Fatalf("Input file %s does not exist.", file)
		}

		fmt.Printf("Running solution for Day: %d, Part: %d, File: %s\n", day, part, fileType)
		switch day {
		case 1:
			if part == 1 {
				day01.Part1(file)
			} else {
				day01.Part2(file)
			}
		case 2:
			if part == 1 {
				day02.Part1(file)
			} else {
				day02.Part2(file)
			}
		case 3:
			if part == 1 {
				day03.Part1(file)
			} else {
				day03.Part2(file)
			}
		case 4:
			if part == 1 {
				day04.Part1(file)
			} else {
				day04.Part2(file)
			}
		case 5:
			if part == 1 {
				day05.Part1(file)
			} else {
				day05.Part2(file)
			}
		case 6:
			if part == 1 {
				day06.Part1(file)
			} else {
				day06.Part2(file)
			}
		case 7:
			if part == 1 {
				day07.Part1(file)
			} else {
				day07.Part2(file)
			}
		case 8:
			if part == 1 {
				day08.Part1(file)
			} else {
				day08.Part2(file)
			}
		case 9:
			if part == 1 {
				day09.Part1(file)
			} else {
				day09.Part2(file)
			}
		default:
			log.Fatalf("Solution for day %d is not implemented yet.", day)
		}
	},
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func init() {
	rootCmd.AddCommand(runnerCmd)
	runnerCmd.Flags().IntVar(&day, "day", 0, "Specify the day (1-25)")
	runnerCmd.Flags().IntVar(&part, "part", 0, "Specify the part (1 or 2)")
	runnerCmd.Flags().
		StringVar(&fileType, "file", "", "Specify a filename type (example or input)")
}
