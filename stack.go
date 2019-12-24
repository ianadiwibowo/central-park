package data_structure

import (
	"fmt"
)

type Stack struct {
	Storage []int
}

// NewStack creates a new empty stack
func NewStack() *Stack {
	return &Stack{
		Storage: []int{},
	}
}

// Push puts the value to the top of the stack
func (s *Stack) Push(value int) {
	s.Storage = append(s.Storage, value)
}

// Pop returns the top value and remove it from the stack
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	n := len(s.Storage) - 1
	value := s.Storage[n]
	s.Storage[n] = 0
	s.Storage = s.Storage[:n]

	return value
}

// Peek returns the top value without removing it from the stack
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}

	n := len(s.Storage) - 1
	value := s.Storage[n]

	return value
}

// IsEmpty tells whether the stack is empty (true) or not (false)
func (s *Stack) IsEmpty() bool {
	return len(s.Storage) == 0
}

// Print returns the human-readable format of the stack
func (s *Stack) Print() string {
	return fmt.Sprintf("%v", s.Storage)
}
