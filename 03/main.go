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

func largestJolt(bank []int) int {
	// You always want to use the largest digit you can for the 10s place. Using
	// a lower digit will always result in a smaller number. However, you cannot
	// use the last number because you would have no 1s place digit.
	firstIndex := firstHighestIndex(bank[:len(bank) - 1])

	// Out of the remaining digits, we always want the biggest one to make the
	// largest overall number
	remainingStart := firstIndex + 1
	secondIndex := remainingStart + firstHighestIndex(bank[remainingStart:])

	return 10 * bank[firstIndex] + bank[secondIndex]
}

