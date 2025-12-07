package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type SimulationResult struct{
	// The number of times a beam split on a splitter
	splits uint

	// The number of timelines generated
	timelines uint
}

/*
Simulate a manifold and calculate information about it

Beams start from any 'S' in the first row.

The manifold cannot consecutive splitters '^', or have a splitter at the edge of
a row.
*/
func simulate(manifold [][]byte) SimulationResult {
	beams := make([]uint, len(manifold[0]))
	for i, char := range manifold[0] {
		if char == 'S' {
			beams[i] = 1
		}
	}

	splitCount := uint(0)
	for row := range manifold {
		for col, char := range manifold[row] {
			if char == '^' && beams[col] > 0 {
				// Skip the bounds check because because manifold doesn't have a
				// splitter at the edge
				beams[col - 1] += beams[col]
				beams[col + 1] += beams[col]
				beams[col] = 0

				splitCount += 1
			}
		}
	}

	totalBeams := uint(0)
	for _, count := range beams {
		totalBeams += count
	}

	return SimulationResult{
		splitCount,
		totalBeams,
	}
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

	results := simulate(manifold)
	fmt.Printf("Split count: %d\n", results.splits)
	fmt.Printf("Timelines: %d\n", results.timelines)
}

