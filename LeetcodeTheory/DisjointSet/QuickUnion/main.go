package main

import "fmt"

/**
Quick Union
Stores the parent of every element

Quick Find: finding is fast, union is slow.
Quick Union: union is faster, but finding may require following parent links.
*/

func main() {
	ds := NewQuickUnion(6)
	ds.Union(0, 1)
	ds.Union(1, 2)
	ds.Union(3, 4)
	fmt.Println(ds.Connected(0, 2)) // true
	fmt.Println(ds.Connected(0, 4)) // false
	ds.Union(2, 4)
	fmt.Println(ds.Connected(0, 4)) // true
	fmt.Println(ds.parents)         // [0 0 0 0 0 5]
}

type QuickUnion struct {
	parents []int
}

func NewQuickUnion(len int) *QuickUnion {
	parents := make([]int, len)
	for i, _ := range parents {
		parents[i] = i
	}
	return &QuickUnion{
		parents: parents,
	}
}

// Find follows parent links until it reaches the root.
func (q *QuickUnion) Find(element int) int {
	for element != q.parents[element] {
		element = q.parents[element]
	}

	return element
}
func (q *QuickUnion) Union(a, b int) {
	rootA := q.Find(a)
	rootB := q.Find(b)
	if rootA == rootB {
		return
	}

	q.parents[rootB] = rootA
}

// Connected checks whether two elements have the same root.
func (q *QuickUnion) Connected(a, b int) bool {
	return q.Find(a) == q.Find(b)
}
