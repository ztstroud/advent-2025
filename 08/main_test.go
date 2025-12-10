package main

import (
	"reflect"
	"testing"
)

func TestNShortestEdges(t *testing.T) {
	ps := []Point{
		{ 0, 0, 0 },
		{ 1, 0, 0 },
		{ 0, 2, 0 },
		{ 0, 0, 3 },
	}

	shortest := nShortestEdges(ps, 3)
	// This is sensitive to the order of the result, even though that actually
	// doesn't matter and could change
	expected := []Edge{
		{ 1, 2, 5 },
		{ 0, 1, 1 },
		{ 0, 2, 4 },
	}

	if !reflect.DeepEqual(shortest, expected) {
		t.Errorf("Expected %v to be %v", shortest, expected)
	}
}

func TestNLargestGroups(t *testing.T) {
	set := NewDisjointSet(10)

	set.Merge(0, 1)
	set.Merge(2, 3)
	set.Merge(0, 3)

	set.Merge(4, 5)
	set.Merge(5, 6)

	set.Merge(7, 8)

	largest := nLargestGroups(set, 3)
	// This is sensitive to the order of the result, even though that actually
	// doesn't matter and could change
	expected := []uint{ 7, 0, 4 }

	if !reflect.DeepEqual(largest, expected) {
		t.Errorf("Expected %v to be %v", largest, expected)
	}
}

