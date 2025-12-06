package main

import (
	"reflect"
	"testing"
)

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

