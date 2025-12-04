package main

func firstHighestIndex(vs []int) int {
	if len(vs) == 0 {
		return -1
	}

	highest := vs[0]
	for _, v := range vs {
		highest = max(highest, v)
	}

	for i, v := range vs {
		if v == highest {
			return i
		}
	}

	panic("highest value in slice is no longer present")
}

