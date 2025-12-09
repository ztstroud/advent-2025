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

