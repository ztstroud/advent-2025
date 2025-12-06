package main

import "slices"

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

