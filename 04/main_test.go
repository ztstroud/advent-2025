package main

import "testing"

func TestCountAccessibleRolls(t *testing.T) {
	grid := [][]byte{
		[]byte("..@@.@@@@."),
		[]byte("@@@.@.@.@@"),
		[]byte("@@@@@.@.@@"),
		[]byte("@.@@@@..@."),
		[]byte("@@.@@@@.@@"),
		[]byte(".@@@@@@@.@"),
		[]byte(".@.@.@.@@@"),
		[]byte("@.@@@.@@@@"),
		[]byte(".@@@@@@@@."),
		[]byte("@.@.@@@.@."),
	}

	count := countAccessibleRolls(grid)
	expected := 13

	if count != expected {
		t.Errorf("Expected %d to be %d", count, expected)
	}
}

func TestCountAccessibleRollsWithRemoval(t *testing.T) {
	grid := [][]byte{
		[]byte("..@@.@@@@."),
		[]byte("@@@.@.@.@@"),
		[]byte("@@@@@.@.@@"),
		[]byte("@.@@@@..@."),
		[]byte("@@.@@@@.@@"),
		[]byte(".@@@@@@@.@"),
		[]byte(".@.@.@.@@@"),
		[]byte("@.@@@.@@@@"),
		[]byte(".@@@@@@@@."),
		[]byte("@.@.@@@.@."),
	}

	count := countAccessibleRollsWithRemoval(grid)
	expected := 43

	if count != expected {
		t.Errorf("Expected %d to be %d", count, expected)
	}
}

