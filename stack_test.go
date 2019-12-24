package data_structure_test

import (
	"testing"

	data_structure "github.com/ianadiwibowo/central-park"
)

func TestStackNewStack(t *testing.T) {
	s := data_structure.NewStack()

	if s.ToString() != "[]" {
		t.Errorf("Expected: []. Got: %v", s.ToString())
	}
}

func TestStackPush(t *testing.T) {
	s := data_structure.NewStack()
	s.Push(5)
	s.Push(3)

	if s.ToString() != "[5 3]" {
		t.Errorf("Expected: [5 3], Got: %v", s.ToString())
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

	if s.ToString() != "[5]" {
		t.Errorf("Expected: [5], Got: %v", s.ToString())
	}
}

func TestStackPopOnEmptyStack(t *testing.T) {
	s := data_structure.NewStack()
	v := s.Pop()

	if v != -1 {
		t.Errorf("Expected: -1, Got: %v", v)
	}

	if s.ToString() != "[]" {
		t.Errorf("Expected: [], Got: %v", s.ToString())
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

	if s.ToString() != "[5 3]" {
		t.Errorf("Expected: [5 3], Got: %v", s.ToString())
	}
}

func TestStackPeekOnEmptyStack(t *testing.T) {
	s := data_structure.NewStack()
	v := s.Peek()

	if v != -1 {
		t.Errorf("Expected: -1, Got: %v", v)
	}

	if s.ToString() != "[]" {
		t.Errorf("Expected: [], Got: %v", s.ToString())
	}
}

func TestStackIsEmptyOnEmptyStack(t *testing.T) {
	s := data_structure.NewStack()

	if s.IsEmpty() == false {
		t.Errorf("Expected: true, Got: %v", s.ToString())
	}
}

func TestStackIsEmptyOnNotEmptyStack(t *testing.T) {
	s := data_structure.NewStack()
	s.Push(5)
	s.Push(3)

	if s.IsEmpty() == true {
		t.Errorf("Expected: false, Got: %v", s.ToString())
	}
}
