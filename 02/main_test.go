package main

import (
	"fmt"
	"testing"
)

func TestNumDigits(t *testing.T) {
	cases := []struct{
		val uint64
		expected int
	}{
		{ 1, 1 },
		{ 9, 1 },
		{ 10, 2 },
		{ 99, 2 },
		{ 100, 3 },
		{ 999, 3 },
		{ 1000, 4 },
		{ 9999, 4 },
		
		// The upper bound of the puzzle input
		{ 1000000000, 10 },
		{ 9999999999, 10 },
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d", c.val), func(t *testing.T) {
			digits := numDigits(c.val)
			
			if digits != c.expected {
				t.Errorf("Expected %d to be %d", digits, c.expected)
			}
		})
	}
}

