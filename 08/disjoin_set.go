package main

type DisjointSet struct{
	parents []uint
	size []uint
}

func NewDisjointSet(n uint) DisjointSet {
	parents := make([]uint, n)
	size := make([]uint, n)

	for i := range n {
		parents[i] = i
		size[i] = 1
	}

	return DisjointSet{ parents, size }
}

func (set DisjointSet) Find(i uint) uint {
	if set.parents[i] == i {
		return i
	}

	root := set.Find(set.parents[i])
	set.parents[i] = root

	return root
}

func (set DisjointSet) Merge(a, b uint) {
	a = set.Find(a)
	b = set.Find(b)

	if a == b {
		return
	}

	if set.size[a] < set.size[b] {
		a, b = b, a
	}

	set.parents[b] = a
	set.size[a] += set.size[b]
}

func (set DisjointSet) Count() uint {
	count := uint(0)
	for i, p := range set.parents {
		if uint(i) == p {
			count += 1
		}
	}

	return count
}

func (set DisjointSet) Size(i uint) uint {
	return set.size[set.Find(i)]
}

