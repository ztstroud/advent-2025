package main

/*
Compares to values. Returns:
- A positive value if a > b
- Zero if a == b
- A negative value if a < b
*/
type Cmp[T any] func(a, b T) int

type Heap[T any] struct{
	data []T
	cmp Cmp[T]
}

/*
Create a new Heap with the given data

The given data is copied and not modified
*/
func NewHeap[T any](n int, cmp func(a, b T) int) Heap[T] {
	heap := Heap[T]{
		make([]T, 0, n),
		cmp,
	}

	return heap
}

func (heap Heap[T]) bubbleUp(i int) {
	pi := (i - 1) / 2
	for heap.cmp(heap.data[i], heap.data[pi]) > 0 {
		tmp := heap.data[i]
		heap.data[i] = heap.data[pi]
		heap.data[pi] = tmp

		i = pi
		pi = (i - 1) / 2
	}
}

func (heap Heap[T]) bubbleDown(i int) {
	for i < len(heap.data) {

		lci := i * 2 + 1
		rci := lci + 1

		if lci >= len(heap.data) {
			break
		}

		if rci >= len(heap.data) {
			if heap.cmp(heap.data[lci], heap.data[i]) > 0 {
				tmp := heap.data[i]
				heap.data[i] = heap.data[lci]
				heap.data[lci] = tmp

				i = lci
			} else {
				break
			}
		} else if heap.cmp(heap.data[lci], heap.data[rci]) > 0 {
			if heap.cmp(heap.data[lci], heap.data[i]) > 0 {
				tmp := heap.data[i]
				heap.data[i] = heap.data[lci]
				heap.data[lci] = tmp

				i = lci
			} else {
				break
			}
		} else {
			if heap.cmp(heap.data[rci], heap.data[i]) > 0 {
				tmp := heap.data[i]
				heap.data[i] = heap.data[rci]
				heap.data[rci] = tmp

				i = rci
			} else {
				break
			}
		}
	}
}

func (heap *Heap[T]) Len() int {
	return len(heap.data)
}

func (heap *Heap[T]) Peek() T {
	return heap.data[0]
}

func (heap *Heap[T]) Insert(new T) {
	if len(heap.data) >= cap(heap.data) {
		panic("Cannot add more to the heap!")
	}

	heap.data = append(heap.data, new)
	heap.bubbleUp(len(heap.data) - 1)
}

func (heap *Heap[T]) Replace(new T) {
	heap.data[0] = new
	heap.bubbleDown(0)
}

