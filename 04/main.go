package main

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

