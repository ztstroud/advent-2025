package main

import (
	"reflect"
	"testing"
)

func intMaxCmp(a, b int) int {
	return a - b;
}

func TestNewHeap(t *testing.T) {
	capacity := 10
	heap := NewHeap(capacity, intMaxCmp)

	if cap(heap.data) != 10 {
		t.Errorf(
			"Heap had wrong capacity, expected %d to be %d",
			cap(heap.data),
			capacity,
		)
	}
}

func TestHeapInsert(t *testing.T) {
	data := []int{ 3, 5, 1, 6, 4, 7, 2 }
	heap := NewHeap(len(data), intMaxCmp)

	for _, e := range data {
		heap.Insert(e)
	}

	expected := []int{ 7, 5, 6, 3, 4, 1, 2 }

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v to be %v", heap.data, expected)
	}
}

func TestHeapPeek(t *testing.T) {
	data := []int{ 3, 5, 1, 6, 4, 7, 2 }
	heap := NewHeap(len(data), intMaxCmp)

	for _, e := range data {
		heap.Insert(e)
	}

	peeked := heap.Peek()
	expected := 7

	if peeked != expected {
		t.Errorf("Expected %v to be %v", peeked, expected)
	}
}

func TestHeapReplace(t *testing.T) {
	data := []int{ 3, 5, 1, 6, 4, 7, 2 }
	heap := NewHeap(len(data), intMaxCmp)

	for _, e := range data {
		heap.Insert(e)
	}

	heap.Replace(0)
	expected := []int{ 6, 5, 2, 3, 4, 1, 0 }

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v to be %v", heap.data, expected)
	}
}

