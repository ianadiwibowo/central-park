package data_structure_test

import (
	"testing"

	data_structure "github.com/ianadiwibowo/central-park"
)

func TestStackNewStack(t *testing.T) {
	s := data_structure.NewStack()

	if s.Print() != "[]" {
		t.Errorf("Expected: []. Got: %v", s.Print())
	}
}

func TestStackPush(t *testing.T) {
	s := data_structure.NewStack()
	s.Push(5)
	s.Push(3)

	if s.Print() != "[5 3]" {
		t.Errorf("Expected: [5 3], Got: %v", s.Print())
	}
}

func TestStackPop(t *testing.T) {
	s := data_structure.NewStack()
	s.Push(5)
	s.Push(3)
	v := s.Pop()

	if v != 3 {
		t.Errorf("Expected: 3, Got: %v", v)
	}

	if s.Print() != "[5]" {
		t.Errorf("Expected: [5], Got: %v", s.Print())
	}
}

func TestStackPopOnEmptyStack(t *testing.T) {
	s := data_structure.NewStack()
	v := s.Pop()

	if v != -1 {
		t.Errorf("Expected: -1, Got: %v", v)
	}

	if s.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", s.Print())
	}
}

func TestStackPeek(t *testing.T) {
	s := data_structure.NewStack()
	s.Push(5)
	s.Push(3)
	v := s.Peek()

	if v != 3 {
		t.Errorf("Expected: 3, Got: %v", v)
	}

	if s.Print() != "[5 3]" {
		t.Errorf("Expected: [5 3], Got: %v", s.Print())
	}
}

func TestStackPeekOnEmptyStack(t *testing.T) {
	s := data_structure.NewStack()
	v := s.Peek()

	if v != -1 {
		t.Errorf("Expected: -1, Got: %v", v)
	}

	if s.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", s.Print())
	}
}

func TestStackIsEmptyOnEmptyStack(t *testing.T) {
	s := data_structure.NewStack()

	if s.IsEmpty() == false {
		t.Errorf("Expected: true, Got: %v", s.Print())
	}
}

func TestStackIsEmptyOnNotEmptyStack(t *testing.T) {
	s := data_structure.NewStack()
	s.Push(5)
	s.Push(3)

	if s.IsEmpty() == true {
		t.Errorf("Expected: false, Got: %v", s.Print())
	}
}
