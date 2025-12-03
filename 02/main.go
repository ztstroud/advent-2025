package main

import "math"

func numDigits(val uint64) int {
	return int(math.Log10(float64(val))) + 1
}

