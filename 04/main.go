package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const EMPTY = '.'
const ROLL = '@'
const ACCESSIBLE = 'x'

func markAccessibleRolls(grid [][]byte) {
	for y := range grid {
		for x := range grid {
			if grid[y][x] == EMPTY {
				continue
			}

			surroundingCount := 0

			if y > 0 {
				if x > 0 && grid[y - 1][x - 1] != EMPTY {
					surroundingCount += 1
				}

				if grid[y - 1][x] != EMPTY {
					surroundingCount += 1
				}

				if x < len(grid[y]) - 1 && grid[y - 1][x + 1] != EMPTY {
					surroundingCount += 1
				}
			}

			if x > 0 && grid[y][x - 1] != EMPTY {
				surroundingCount += 1
			}

			if x < len(grid[y]) - 1 && grid[y][x + 1] != EMPTY {
				surroundingCount += 1
			}

			if y < len(grid) - 1 {
				if x > 0 && grid[y + 1][x - 1] != EMPTY {
					surroundingCount += 1
				}

				if grid[y + 1][x] != EMPTY {
					surroundingCount += 1
				}

				if x < len(grid[y]) - 1 && grid[y + 1][x + 1] != EMPTY {
					surroundingCount += 1
				}
			}

			if surroundingCount < 4 {
				grid[y][x] = ACCESSIBLE
			}
		}
	}
}

func countEqual(grid [][]byte, query byte) int {
	count := 0
	for y := range grid {
		for x := range grid {
			if grid[y][x] == query {
				count += 1
			}
		}
	}

	return count
}

func removeEqual(grid [][]byte, query byte) {
	for y := range grid {
		for x := range grid {
			if grid[y][x] == query {
				grid[y][x] = EMPTY
			}
		}
	}
}

func countAccessibleRolls(grid [][]byte) int {
	markAccessibleRolls(grid)
	return countEqual(grid, ACCESSIBLE)
}

func countAccessibleRollsWithRemoval(grid [][]byte) int {
	count := 0
	for {
		markAccessibleRolls(grid)
		newlyAccessible := countEqual(grid, ACCESSIBLE)

		if newlyAccessible == 0 {
			break
		}

		count += newlyAccessible
		removeEqual(grid, ACCESSIBLE)
	}
	return count
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

	count := countAccessibleRollsWithRemoval(grid)
	fmt.Printf("Accessible rolls with removal: %d\n", count)
}

