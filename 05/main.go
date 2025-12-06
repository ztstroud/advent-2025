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

