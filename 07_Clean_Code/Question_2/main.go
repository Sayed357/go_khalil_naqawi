package main

import (
	"fmt"
)

type Set struct {
	data map[int]struct{}
}

// NewSet init a new set
func NewSet() *Set {
	return &Set{
		data: make(map[int]struct{}),
	}
}

// Add, adds an integer to the set
func (s *Set) Add(value int) {
	s.data[value] = struct{}{}
}

// Get returns all elements of the set as slice
func (s *Set) Get() []int {
	elements := make([]int, 0, len(s.data))
	for key := range s.data {
		elements = append(elements, key)
	}
	return elements
}

// Remove: deletes an integer from the set.
func (s *Set) Remove(value int) {
	delete(s.data, value)
}

func main() {
	set := NewSet()

	set.Add(1)
	set.Add(2)
	set.Add(1)
	set.Add(3)
	set.Add(4)
	set.Add(5)

	set.Remove(100) //remove a non existand element.

	fmt.Println(set.Get())
}
