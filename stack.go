package stack

import "fmt"

type Stack struct {
	storage []int
}

// New creates a new empty stack
func New() *Stack {
	return &Stack{
		storage: []int{},
	}
}

// Push puts the value to the top of the stack
func (s *Stack) Push(value int) {
	s.storage = append(s.storage, value)
}

// Pop returns the top value and remove it from the stack
func (s *Stack) Pop() int {
	if len(s.storage) == 0 {
		return -1
	}

	n := len(s.storage) - 1
	value := s.storage[n]
	s.storage[n] = 0
	s.storage = s.storage[:n]

	return value
}

// ToString returns the human-readable format of the stack
func (s *Stack) ToString() string {
	return fmt.Sprintf("%v", s.storage)
}
