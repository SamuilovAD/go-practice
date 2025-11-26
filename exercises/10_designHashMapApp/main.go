package main

import "fmt"

/*
*
Task:
Implement a custom hash map
*/
func main() {
	m := NewMyHashMap()

	m.Put(1, 10)
	m.Put(2, 20)
	m.Put(17, 170) // Will likely go to same bucket as 1 (depending on capacity)

	fmt.Println("Get(1):", m.Get(1))   // 10
	fmt.Println("Get(2):", m.Get(2))   // 20
	fmt.Println("Get(17):", m.Get(17)) // 170
	fmt.Println("Get(3):", m.Get(3))   // -1 (not found)

	m.Remove(2)
	fmt.Println("After Remove(2), Get(2):", m.Get(2)) // -1
}

const (
	defaultCapacity = 16
	loadFactor      = 0.75
)

// entry represents a single key-value pair in a bucket.
type entry struct {
	key   int
	value int
	next  *entry
}

// MyHashMap is a simple hash map implementation with separate chaining.
type MyHashMap struct {
	buckets []*entry
	size    int
}

// NewMyHashMap creates a new hash map with default capacity.
func NewMyHashMap() *MyHashMap {
	return &MyHashMap{
		buckets: make([]*entry, defaultCapacity),
	}
}

// hash calculates bucket index for a given key.
func (m *MyHashMap) hash(key int) int {
	if len(m.buckets) == 0 {
		return 0
	}
	if key < 0 {
		key = -key
	}
	return key % len(m.buckets)
}

// Put inserts or updates the value by key.
func (m *MyHashMap) Put(key int, value int) {
	// Resize if load factor exceeded
	if float64(m.size+1) > float64(len(m.buckets))*loadFactor {
		m.resize()
	}

	index := m.hash(key)
	head := m.buckets[index]

	// Check if key already exists in the chain
	for e := head; e != nil; e = e.next {
		if e.key == key {
			e.value = value
			return
		}
	}

	// Insert new entry at the head of the chain
	newEntry := &entry{
		key:   key,
		value: value,
		next:  head,
	}
	m.buckets[index] = newEntry
	m.size++
}

// Get returns value by key or -1 if key is not found.
func (m *MyHashMap) Get(key int) int {
	index := m.hash(key)
	for e := m.buckets[index]; e != nil; e = e.next {
		if e.key == key {
			return e.value
		}
	}
	// Convention: -1 means "not found"
	return -1
}

// Remove deletes key from the map if it exists.
func (m *MyHashMap) Remove(key int) {
	index := m.hash(key)
	current := m.buckets[index]
	var prev *entry

	for current != nil {
		if current.key == key {
			if prev == nil {
				// Remove first element in chain
				m.buckets[index] = current.next
			} else {
				prev.next = current.next
			}
			m.size--
			return
		}
		prev = current
		current = current.next
	}
}

// resize doubles buckets count and rehashes all entries.
func (m *MyHashMap) resize() {
	newCapacity := len(m.buckets) * 2
	if newCapacity == 0 {
		newCapacity = defaultCapacity
	}

	newBuckets := make([]*entry, newCapacity)

	// Rehash all existing entries
	for _, head := range m.buckets {
		for e := head; e != nil; {
			next := e.next // Save next pointer

			// Calculate new index
			key := e.key
			if key < 0 {
				key = -key
			}
			index := key % newCapacity

			// Insert at new head
			e.next = newBuckets[index]
			newBuckets[index] = e

			e = next
		}
	}

	m.buckets = newBuckets
}
