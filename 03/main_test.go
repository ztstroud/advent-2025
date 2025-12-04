package main

import "testing"

func TestFirstHighestIndex(t *testing.T) {
	cases := []struct{
		name string
		vs []int
		expected int
	}{
		{
			"Empty",
			[]int{},
			-1,
		},
		{
			"TwoInstances",
			[]int{ 0, 0, 0, 1, 0, 0 },
			3,
		},
		{
			"TwoInstances",
			[]int{ 0, 1, 0, 0, 1, 0 },
			1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := firstHighestIndex(c.vs)
			
			if actual != c.expected {
				t.Errorf("Expected %d to be %d", actual, c.expected)
			}
		})
	}
}

func TestLargestJolt(t *testing.T) {
	cases := []struct {
		name string
		bank []int
		expected int
	}{
		{
			"Given1",
			[]int{ 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1 },
			98,
		},
		{
			"Given2",
			[]int{ 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9 },
			89,
		},
		{
			"Given3",
			[]int{ 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8 },
			78,
		},
		{
			"Given4",
			[]int{ 8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1 },
			92,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := largestJolt(c.bank)
			
			if actual != c.expected {
				t.Errorf("Expected %d to be %d", actual, c.expected)
			}
		})
	}
}

