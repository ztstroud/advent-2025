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

