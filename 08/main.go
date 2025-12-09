package main

type Point struct{ x, y, z int }

type Edge struct{ start, end, distSquared int }

func cmpEdges(a, b Edge) int {
	return a.distSquared - b.distSquared
}

func squareDist(a, b Point) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z

	return dx * dx + dy * dy + dz * dz
}

func nShortestEdges(ps []Point, n int) []Edge {
	heap := NewHeap(n, cmpEdges)

	for end := range ps {
		for start := 0; start < end; start += 1 {
			edge := Edge{ start, end, squareDist(ps[start], ps[end]) }

			if heap.Len() < n {
				heap.Insert(edge)
			} else if cmpEdges(heap.Peek(), edge) > 0 {
				heap.Replace(edge)
			}
		}
	}

	return heap.data
}

