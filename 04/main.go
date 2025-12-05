package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const EMPTY = '.'
const ROLL = '@'

func countAccessibleRolls(grid [][]byte) int {
	accessibleCount := 0
	for y := range grid {
		for x := range grid {
			if grid[y][x] == EMPTY {
				continue
			}

			surroundingCount := 0

			if y > 0 {
				if x > 0 && grid[y - 1][x - 1] == ROLL {
					surroundingCount += 1
				}

				if grid[y - 1][x] == ROLL {
					surroundingCount += 1
				}

				if x < len(grid[y]) - 1 && grid[y - 1][x + 1] == ROLL {
					surroundingCount += 1
				}
			}

			if x > 0 && grid[y][x - 1] == ROLL {
				surroundingCount += 1
			}

			if x < len(grid[y]) - 1 && grid[y][x + 1] == ROLL {
				surroundingCount += 1
			}

			if y < len(grid) - 1 {
				if x > 0 && grid[y + 1][x - 1] == ROLL {
					surroundingCount += 1
				}

				if grid[y + 1][x] == ROLL {
					surroundingCount += 1
				}

				if x < len(grid[y]) - 1 && grid[y + 1][x + 1] == ROLL {
					surroundingCount += 1
				}
			}

			if surroundingCount < 4 {
				accessibleCount += 1
			}
		}
	}

	return accessibleCount
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must provide an input file\n")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not read file: %s\n%v\n", path, err)
	}
	defer file.Close()

	grid := make([][]byte, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]byte, len(scanner.Bytes()))
		copy(row, scanner.Bytes())

		grid = append(grid, row)
	}

	count := countAccessibleRolls(grid)
	fmt.Printf("Accessible rolls: %d\n", count)
}

