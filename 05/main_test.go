package main

import (
	"reflect"
	"testing"
)

func TestMergeRanges(t *testing.T) {
    cases := []struct{
        name string
		spans []Span
		expected []Span
    }{
		{
			"SameStart",
			[]Span{
				{ 1, 5 },
				{ 1, 10 },
			},
			[]Span{
				{ 1, 10 },
			},
		},
		{
			"DifferentStart",
			[]Span{
				{ 1, 5 },
				{ 3, 10 },
			},
			[]Span{
				{ 1, 10 },
			},
		},
		{
			"EndsOverlap",
			[]Span{
				{ 1, 5 },
				{ 5, 10 },
			},
			[]Span{
				{ 1, 10 },
			},
		},
		{
			"EndsMeet",
			[]Span{
				{ 1, 5 },
				{ 6, 10 },
			},
			[]Span{
				{ 1, 10 },
			},
		},
		{
			"Separate",
			[]Span{
				{ 1, 4 },
				{ 6, 10 },
			},
			[]Span{
				{ 1, 4 },
				{ 6, 10 },
			},
		},
		{
			"EndsSorted",
			[]Span{
				{ 6, 10 },
				{ 1, 4 },
			},
			[]Span{
				{ 1, 4 },
				{ 6, 10 },
			},
		},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
			merged := mergeSpans(c.spans)

			if !reflect.DeepEqual(merged, c.expected) {
				t.Errorf("Expected %v to be %v", merged, c.expected)
			}
        })
    }
}

