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
	if len(s.Storage) == 0 {
		return -1
	}

	n := len(s.Storage) - 1
	value := s.Storage[n]
	s.Storage[n] = 0
	s.Storage = s.Storage[:n]

	return value
}

// Top returns the top value without removing it from the stack
func (s *Stack) Top() int {
	if len(s.Storage) == 0 {
		return -1
	}

	n := len(s.Storage) - 1
	value := s.Storage[n]

	return value
}

// ToString returns the human-readable format of the stack
func (s *Stack) ToString() string {
	return fmt.Sprintf("%v", s.Storage)
}
