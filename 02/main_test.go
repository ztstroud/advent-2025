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

func TestCountRepeatedFullRange(t *testing.T) {
	cases := []struct {
		name string
		lower, upper uint64
		expected uint64
	}{
		{
			"SingleDigit",
			1, 9,
			0,
		},
		{
			"DoubleDigit",
			10, 99,
			495,
		},
		{
			"TripleDigit",
			100, 999,
			0,
		},
		{
			"QuadrupleDigit",
			1000, 9999,
			495405,
		},
		{
			"SingleToQuadrupleDigit",
			1, 9999,
			495900,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			digits := countRepeated(c.lower, c.upper)
			
			if digits != c.expected {
				t.Errorf("Expected %d to be %d", digits, c.expected)
			}
		})
	}
}

