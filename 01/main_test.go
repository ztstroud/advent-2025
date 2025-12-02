package main

import "testing"

func TestCountZerosPositive(t *testing.T) {
	seq := []int{ 25, 100, 75 }

	count := countZeros(0, 100, seq)
	expected := 1

	if count != expected {
		t.Errorf("Expected %v to be %v", count, expected)
	}
}

func TestCountZerosNegative(t *testing.T) {
	seq := []int{ -25, -100, -75 }

	count := countZeros(0, 100, seq)
	expected := 1

	if count != expected {
		t.Errorf("Expected %v to be %v", count, expected)
	}
}

func TestParseTurnValid(t *testing.T) {
	cases := []struct{
		input string
		expected int
	}{
		{ "R45", 45 },
		{ "L45", -45 },
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			parsed, err := parseTurn(c.input)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if parsed != c.expected {
				t.Errorf("Expected %v to be %v", parsed, c.expected)
			}
		})
	}
}

func TestParseTurnInvalid(t *testing.T) {
	cases := []struct{
		name string
		input string
	}{
		{ "Invalid direction", "Z45" },
		{ "Invalid number", "RABC" },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := parseTurn(c.input)

			if err == nil {
				t.Errorf("Expected an error, but not received")
			}
		})
	}
}

