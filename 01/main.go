package main

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

