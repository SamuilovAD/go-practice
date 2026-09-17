package main

import "fmt"

/*
Union-find Constructor: Time O(N)
Find: Time O(1)
Union: Time O(N)
Connected: Time O(N)
*/
type QuickFind struct {
	group []int
}

func NewQuickFind(size int) *QuickFind {
	group := make([]int, size)
	// Initially, every element belongs to its own group.
	for i := range group {
		group[i] = i
	}

	return &QuickFind{group: group}
}

func (q *QuickFind) Find(element int) int {
	return q.group[element]
}

func (q *QuickFind) Connected(a, b int) bool {
	return q.Find(a) == q.Find(b)
}

func (q *QuickFind) Union(a, b int) {
	groupA := q.Find(a)
	groupB := q.Find(b)
	if groupA == groupB {
		return
	}
	// Replace all occurrences of groupB with groupA.
	for i := range q.group {
		if q.group[i] == groupB {
			q.group[i] = groupA
		}
	}
}

func main() {
	ds := NewQuickFind(6)
	ds.Union(0, 1)
	ds.Union(1, 2)
	ds.Union(3, 4)
	fmt.Println(ds.Connected(0, 2)) // true
	fmt.Println(ds.Connected(0, 4)) // false
	ds.Union(2, 4)
	fmt.Println(ds.Connected(0, 4)) // true
	fmt.Println(ds.group)           // [0 0 0 0 0 5]
}
