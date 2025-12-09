package main

import "testing"

func TestDisjointSet(t *testing.T) {
	set := NewDisjointSet(5)

	set.Merge(0, 1)
	set.Merge(2, 3)
	set.Merge(3, 4)
	set.Merge(3, 1)

	expectedParent := uint(2)

	for i := range uint(5) {
		found := set.Find(i)
		if found != expectedParent {
			t.Errorf("All sets should find %d, found %d", expectedParent, found)
		}
	}

	if set.size[expectedParent] != 5 {
		t.Errorf(
			"Root size should be %d, got %d",
			expectedParent,
			set.size[expectedParent],
		)
	}
}

func TestDisjointSetCount(t *testing.T) {
	set := NewDisjointSet(5)

	if set.Count() != 5 {
		t.Errorf(
			"Initial size should be same as given size (got %d, expected 5)",
			set.Count(),
		)
	}

	set.Merge(0, 1)
	set.Merge(2, 3)

	if set.Count() != 3 {
		t.Errorf(
			"Expected %d to be 3",
			set.Count(),
		)
	}
}

