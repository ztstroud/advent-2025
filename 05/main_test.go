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

func TestInAnySpan(t *testing.T) {
	spans := []Span{
		{ 5, 10 },
		{ 15, 20 },
	}

    cases := []struct{
        name string
		val uint64
		expected bool
    }{
		{
			"BelowBoth",
			2,
			false,
		},
		{
			"InLower",
			7,
			true,
		},
		{
			"InBetween",
			12,
			false,
		},
		{
			"InUpper",
			17,
			true,
		},
		{
			"AboveBoth",
			25,
			false,
		},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
			result := inAnySpan(c.val, spans)

			if result != c.expected {
				t.Errorf("Expected %t to be %t", result, c.expected)
			}
        })
    }
}

