# Advent Of Code 2024

Created using go version go1.23.1 linux/amd64

## Running solutions

This repository uses [Cobra](https://github.com/spf13/cobra) to setup CLI interaction

Usage:

```
$ go run aoc runner --help
Run a specific Advent of Code solution for a given day, part, and input file.

Usage:
  aoc runner [flags]

Flags:
      --day int       Specify the day (1-25)
      --file string   Specify a filename type (example or input)
  -h, --help          help for runner
      --part int      Specify the part (1 or 2)
```

Example:

```
go run aoc runner --day 1 --part 1 --file input
```
