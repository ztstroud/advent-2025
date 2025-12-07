package main

import "testing"

func TestCountSplits(t *testing.T) {
	manifold := [][]byte {
		[]byte("...S..."),
		[]byte("......."),
		[]byte("...^..."),
		[]byte("......."),
		[]byte("..^.^.."),
		[]byte("......."),
		[]byte(".^..^.."),
	}

	actual := countSplits(manifold)
	expected := uint(4)

	if actual != expected {
		t.Errorf("Expected %d to be %d", actual, expected)
	}
}

