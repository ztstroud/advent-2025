package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Span struct{ start, end uint64 }

/*
Sort and merge spans to remove duplicates

I don't want to modify the input array, so I map a copy. However, I also want to
avoid unneeded allocations so I execute the merge in place. This means modifying
the copy here, then taking a sub-slice out of the allocated space.
*/
func mergeSpans(spans []Span) []Span {
	merged := make([]Span, len(spans))
	copy(merged, spans)

	slices.SortFunc(merged, func(a, b Span) int {
		if a.start < b.start {
			return -1
		} else if a.start > b.start {
			return 1
		} else {
			return 0
		}
	})

	endIndex := 0
	for i := 1; i < len(merged); i += 1 {
		top := &merged[endIndex]
		candidate := &merged[i]

		if top.end + 1 >= candidate.start {
			top.end = max(top.end, candidate.end)
		} else {
			merged[endIndex + 1] = *candidate
			endIndex += 1
		}
	}

	return merged[:endIndex + 1]
}

/*
Check if the given val is in any of the given spans.

The given spans must be sorted by their start value.
*/
func inAnySpan(val uint64, spans []Span) bool {
	start := 0
	end := len(spans)

	for start < end {
		mid := (start + end) / 2

		if val < spans[mid].start {
			end = mid
		} else if val > spans[mid].end {
			start = mid + 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must specify an input file")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read from file: %s\n%v\n", path, err)
	}

	defer file.Close()

	spans := make([]Span, 0)
	vals := make([]uint64, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		spanSrc := strings.Split(text, "-")
		if len(spanSrc) != 2 {
			log.Fatalf("Invalid range: %s\n", text)
		}

		lower, err := strconv.ParseInt(spanSrc[0], 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse lower bound: %v\n", err)
		}

		upper, err := strconv.ParseInt(spanSrc[1], 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse upper bound: %v\n", err)
		}

		spans = append(spans, Span{ uint64(lower), uint64(upper) })
	}

	for scanner.Scan() {
		val, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse value: %v\n", err)
		}

		vals = append(vals, uint64(val))
	}

	merged := mergeSpans(spans)

	count := 0
	for _, val := range vals {
		if inAnySpan(val, merged) {
			count += 1
		}
	}

	fmt.Printf("In range count: %d\n", count)
}

