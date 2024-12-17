# Advent Of Code 2024

## Running solutions

This repository uses [Cobra](https://github.com/spf13/cobra) to setup CLI interaction

```
$ go run aoc runner --help
Run a specific Advent of Code solution for a given day, part, and input file.

Usage:
  aoc runner [flags]

Flags:
      --day int       Specify the day (1-25)
      --file string   Specify a filename type (example or input)
  -h, --help          help for runner
      --part int      Specify the part (1 or 2)go run aoc runner --day <N>, --part <1/2>, --file <example/input>
```

Example:

```
go run aoc runner --day 1 --part 1 --file example
```
