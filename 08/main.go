package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func nLargestGroups(sets DisjointSet, n int) []uint {
	cmp := func(a, b uint) int {
		return int(sets.Size(b) - sets.Size(a))
	}

	heap := NewHeap(n, cmp)
	for i := range uint(len(sets.parents)) {
		if sets.parents[i] == i {
			if heap.Len() < n {
				heap.Insert(i)
			} else if cmp(heap.Peek(), i) > 0 {
				heap.Replace(i)
			}
		}
	}

	return heap.data
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must specify an input file\n")
	}

	if len(os.Args) < 3 {
		log.Fatal("You must specify a pairing size\n")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read from file: %s\n%v\n", path, err)
	}

	defer file.Close()

	points := make([]Point, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		coordSrcs := strings.Split(line, ",")
		if len(coordSrcs) != 3 {
			log.Fatalf("Invalid coord: %s\n", line)
		}

		x, err := strconv.Atoi(coordSrcs[0])
		if err != nil {
			log.Fatalf("Invalid x: %s\n", coordSrcs[0])
		}

		y, err := strconv.Atoi(coordSrcs[1])
		if err != nil {
			log.Fatalf("Invalid y: %s\n", coordSrcs[1])
		}

		z, err := strconv.Atoi(coordSrcs[2])
		if err != nil {
			log.Fatalf("Invalid z: %s\n", coordSrcs[2])
		}

		points = append(points, Point{ x, y, z })
	}

	pairingSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid pairing size: %s\n", os.Args[2])
	}

	groups := NewDisjointSet(uint(len(points)))
	for _, edge := range nShortestEdges(points, pairingSize) {
		groups.Merge(uint(edge.start), uint(edge.end))
	}

	product := uint(1)
	for _, largestGroupIndex := range nLargestGroups(groups, 3) {
		product *= groups.Size(largestGroupIndex)
	}

	fmt.Printf("Group product: %d\n", product)
}

