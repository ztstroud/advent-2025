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

