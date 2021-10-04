package main

import "fmt"

// Stack structure holding items
type Stack struct {
	items []float64
	err   error
}

// Push an item onto the stack
func (s *Stack) Push(item float64) *Stack {
	s.items = append(s.items, item)
	return s
}

// Pop an item onto the stack
func (s *Stack) Pop() float64 {
	l := len(s.items)
	if l == 0 {
		panic(errStackTooSmall)
	}

	item := s.items[l-1]
	s.items = s.items[0 : l-1]

	return item
}

// Dup top stack item
func (s *Stack) Dup() *Stack {
	if s.IsEmpty() {
		return s
	}

	item := s.Pop()
	s.Push(item).Push(item)

	return s
}

// Swap two top most item
func (s *Stack) Swap() *Stack {
	if len(s.items) < 2 {
		return s
	}

	a, b := s.Pop(), s.Pop()
	s.Push(a).Push(b)

	return s
}

// Size returns stack size
func (s *Stack) Size() int {
	return len(s.items)
}

// IsEmpty returns true if stack is empty, false otherwise
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Flush empties the stack
func (s *Stack) Flush() {
	s.items = s.items[0:]
}

func (s *Stack) String() string {
	var val string
	for _, s := range s.items {
		val = fmt.Sprintf("%s%g\n", val, s)
	}

	return val
}
