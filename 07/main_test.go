package main

import "testing"

func TestSimulate(t *testing.T) {
	manifold := [][]byte {
		[]byte("...S..."),
		[]byte("......."),
		[]byte("...^..."),
		[]byte("......."),
		[]byte("..^.^.."),
		[]byte("......."),
		[]byte(".^..^.."),
	}

	actual := simulate(manifold)
	expected := SimulationResult{
		uint(4),
		uint64(5),
	}

	if actual != expected {
		t.Errorf("Expected %d to be %d", actual, expected)
	}
}

