package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
Count the number of times a beam starting at 'S' will split on splitters '^'

The manifold cannot consecutive splitters, or have a splitter at the edge of a
row.
*/
func countSplits(manifold [][]byte) uint {
	beams := make([]bool, len(manifold[0]))
	for i, char := range manifold[0] {
		if char == 'S' {
			beams[i] = true
		}
	}

	splitCount := uint(0)
	for row := range manifold {
		for col, char := range manifold[row] {
			if char == '^' && beams[col] {
				// Skip the bounds check because because manifold doesn't have a
				// splitter at the edge
				beams[col - 1] = true
				beams[col] = false
				beams[col + 1] = true

				splitCount += 1
			}
		}
	}

	return splitCount
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must specify an input file\n")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read from file: %s\n%v\n", path, err)
	}
	defer file.Close()

	manifold := make([][]byte, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]byte, len(scanner.Bytes()))
		copy(row, scanner.Bytes())

		manifold = append(manifold, row)
	}

	splitCount := countSplits(manifold)
	fmt.Printf("Split count: %d\n", splitCount)
}

