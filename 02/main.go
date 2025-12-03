package main

import "math"

func numDigits(val uint64) int {
	return int(math.Log10(float64(val))) + 1
}

func pow10(val int) uint64 {
	return uint64(math.Pow10(val))
}

func countRepeated(lower, upper uint64) uint64 {
	lowerDigits := numDigits(lower)
	upperDigits := numDigits(upper)

	count := uint64(0)
	for digits := lowerDigits; digits <= upperDigits; digits += 1 {
		// Odd number of digits can never have repeats
		if digits % 2 == 1 {
			continue
		}

		halfMagnitude := pow10(digits / 2)

		countStart := pow10(digits / 2 - 1)
		countEnd := halfMagnitude - 1

		for halfVal := countStart; halfVal <= countEnd; halfVal += 1 {
			count += halfVal + halfVal * halfMagnitude
		}
	}

	return count
}

