package main

/*
Count the number of times a beam starting at 'S' will split on splitters '^'

The manifold cannot consecutive splitters, or have a splitter at the edge of a
row.
*/
func countSplits(manifold [][]byte) uint {
	beams := make([]bool, len(manifold[0]))
	for i, char := range manifold[0] {
		if char == 'S' {
			beams[i] = true
		}
	}

	splitCount := uint(0)
	for row := range manifold {
		for col, char := range manifold[row] {
			if char == '^' && beams[col] {
				// Skip the bounds check because because manifold doesn't have a
				// splitter at the edge
				beams[col - 1] = true
				beams[col] = false
				beams[col + 1] = true

				splitCount += 1
			}
		}
	}

	return splitCount
}

