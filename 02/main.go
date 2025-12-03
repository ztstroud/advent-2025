package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

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
		subMagnitude := pow10(digits / 2 - 1)

		// The diff is the how much you would have to add between each
		// subsequent invalid ID
		diff := halfMagnitude + 1

		// The basis is the value of the first invalid ID in the current number
		// of digits
		basis := subMagnitude * diff

		// n is zero based index of the id in the current number of digits
		sumTo := func(n uint64) uint64 {
			return basis * (n + 1) + diff * (n + 1) * n / 2
		}

		countStart := subMagnitude

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

		startIndex := countStart - subMagnitude
		endIndex := countEnd - subMagnitude

		count += sumTo(endIndex) - sumTo(startIndex - 1)
	}

	return count
}

func isMadeOfRepeating(val uint64) bool {
	for i := 1; i < numDigits(val); i += 1 {
		window := pow10(i)
		windowVal := val % window

		slide := val
		windowRepeating := true

		for slide > 0 {
			if slide % window != windowVal {
				windowRepeating = false
				break
			}

			slide /= window
		}

		if windowRepeating {
			return true
		}
	}

	return false
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must provide an input file\n")
	}

	path := os.Args[1]
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read from file: %s\n%v\n", path, err)
	}

	content := strings.TrimSpace(string(bytes))
	sum := uint64(0)

	for _, rangeSrc := range strings.Split(content, ",") {
		limits := strings.Split(rangeSrc, "-")

		if len(limits) != 2 {
			log.Fatalf("Invalid range: %s\n", rangeSrc)
		}

		lower, err := strconv.ParseUint(limits[0], 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse lower bound: %s\n%v\n", limits[0], err)
		}

		upper, err := strconv.ParseUint(limits[1], 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse upper bound: %s\n%v\n", limits[1], err)
		}

		sum += countRepeated(lower, upper)
	}

	fmt.Printf("Sum of invalid IDs: %d\n", sum)
}

