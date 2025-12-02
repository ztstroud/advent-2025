package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func countZeros(initial, size int, seq []int) int {
	count := 0
	for _, val := range seq {
		initial += val

		if initial % size == 0 {
			count += 1
		}
	}

	return count
}

func parseTurn(turn string) (int, error) {
	sign := 0
	switch turn[0] {
	case 'R':
		sign = 1
	case 'L':
		sign = -1
	default:
		return 0, fmt.Errorf("Invalid turn direction in: %s", turn)
	}

	val, err := strconv.Atoi(turn[1:])
	if err != nil {
		return 0, fmt.Errorf("Invalid turn number in: %s", err)
	}

	return sign * val, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must provide an input path\n")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read file: %s\n%v\n", path, err)
	}
	defer file.Close()

	seq := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		turn, err := parseTurn(scanner.Text())
		if err != nil {
			log.Fatalf("Failed to parse turn: %v\n", err)
		}

		seq = append(seq, turn)
	}

	count := countZeros(50, 100, seq)
	fmt.Printf("Zeros: %d\n", count)
}

