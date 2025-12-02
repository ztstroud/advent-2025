package main

import (
	"fmt"
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

