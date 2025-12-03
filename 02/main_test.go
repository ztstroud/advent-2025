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

func TestCountRepeatedLowerBoundToFull(t *testing.T) {
	cases := []struct {
		name string
		lower, upper uint64
		expected uint64
	}{
		{
			"DoubleDigit",
			20, 99,
			484,
		},
		{
			"DoubleDigitOffset",
			35, 99,
			429,
		},
		{
			"QuadrupleDigit",
			2000, 9999,
			480760,
		},
		{
			"QuadrupleDigitOffset",
			6999, 9999,
			256035,
		},
		{
			"SingleToQuadrupleDigit",
			57, 9999,
			495735,
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

func TestCountRepeatedFullToUpperBound(t *testing.T) {
	cases := []struct {
		name string
		lower, upper uint64
		expected uint64
	}{
		{
			"DoubleDigit",
			10, 89,
			396,
		},
		{
			"DoubleDigitOffset",
			10, 75,
			231,
		},
		{
			"QuadrupleDigit",
			1000, 7999,
			314615,
		},
		{
			"QuadrupleDigitOffset",
			1000, 6241,
			186446,
		},
		{
			"SingleToQuadrupleDigit",
			10, 2711,
			31401,
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

func TestCountRepeatedLowerBoundToUpperBound(t *testing.T) {
	cases := []struct {
		name string
		lower, upper uint64
		expected uint64
	}{
		{
			"None",
			56, 64,
			0,
		},
		{
			"Single",
			55, 55,
			55,
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

func TestIsMadeOfRepeating(t *testing.T) {
	cases := []struct {
		name string
		val uint64
		expected bool
	}{
		{
			"None",
			56,
			false,
		},
		{
			"NoneBig",
			7168419,
			false,
		},
		{
			"Almost",
			123133123,
			false,
		},
		{
			"SingleRepeating",
			55,
			true,
		},
		{
			"DoubleRepeating",
			1313,
			true,
		},
		{
			"TripleRepeating",
			741741,
			true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := isMadeOfRepeating(c.val)

			if result != c.expected {
				t.Errorf("Expected %t to be %t", result, c.expected)
			}
		})
	}
}

