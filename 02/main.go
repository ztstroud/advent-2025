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

		// The lower bound doesn't have the same number of digits as this band,
		// it CANNOT impact the lower bound of this band
		if lowerDigits == digits {
			countStart = lower / halfMagnitude

			// We do need to make sure that this specific value will be above
			// the lower bound, and if not bump it up
			if countStart + countStart * halfMagnitude < lower {
				countStart += 1
			}
		}

		// The same but inverse logic applies to the end count
		countEnd := halfMagnitude - 1
		if upperDigits == digits {
			countEnd = upper / halfMagnitude

			if countEnd + countEnd * halfMagnitude > upper {
				countEnd -= 1
			}
		}

		for halfVal := countStart; halfVal <= countEnd; halfVal += 1 {
			count += halfVal + halfVal * halfMagnitude
		}
	}

	return count
}

