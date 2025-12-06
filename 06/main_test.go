package main

import (
	"reflect"
	"testing"
)

func TestSolveEquation(t *testing.T) {
    cases := []struct{
        name string
		eq Equation
		expected uint64
    }{
		{
			"Addition",
			Equation{ ADD, []uint64{ 23, 300 } },
			323,
		},
		{
			"Multiplication",
			Equation{ MUL, []uint64{ 84, 27 } },
			2268,
		},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
			result := solveEquation(c.eq)

			if result != c.expected {
				t.Errorf("Expected %d to be %d", result, c.expected)
			}
        })
    }
}

func TestParseEqations(t *testing.T) {
	result, err := parseEquations([]string{
		"23  84",
		"300 27",
		"+   *",
	})

	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}

	expected := []Equation{
		{
			ADD,
			[]uint64{ 23, 300 },
		},
		{
			MUL,
			[]uint64{ 84, 27 },
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v to be %v", result, expected)
	}
}

