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

/*
Differs from count zeros in that it keeps the value in the range [0, size) at
the end of all iteration.
*/
func countZeroPasses(initial, size int, seq []int) int {
	val := initial

	count := 0
	for _, turn := range seq {
		newVal := val + turn

		// Interestingly, go's math.Abs function is only for float64
		// They really don't seem to want to add it:
		// https://github.com/golang/go/issues/60623
		passes := newVal / size
		if passes < 0 {
			passes *= -1
		}

		// Large numbers and going past size-1 is handled here
		// This includes turning positively to +/-size, as abs(+/-size/size) = 1
		count += passes

		// We do need to specially track turning exactly to zero, or less than
		// zero. This only applies if you were greater than zero to start, as if
		// you were at zero, you are leaving zero and it should not be counted.
		if val > 0 && newVal <= 0 {
			count += 1
		}

		initial %= size
		if initial < 0 {
			initial += size
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

	passes := countZeroPasses(50, 100, seq)
	fmt.Printf("Passes: %d\n", passes)
}

