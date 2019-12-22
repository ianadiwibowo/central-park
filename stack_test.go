package data_structure_test

import (
	"testing"

	data_structure "github.com/ianadiwibowo/central-park"
)

func TestNewStack(t *testing.T) {
	s := data_structure.NewStack()

	if s.ToString() != "[]" {
		t.Errorf("Expected: []. Got: %v", s.ToString())
	}
}

func TestPush(t *testing.T) {
	s := data_structure.NewStack()
	s.Push(5)
	s.Push(3)

	if s.ToString() != "[5 3]" {
		t.Errorf("Expected: [5 3], Got: %v", s.ToString())
	}
}

func TestPop(t *testing.T) {
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

func TestPopOnEmptyStack(t *testing.T) {
	s := data_structure.NewStack()
	v := s.Pop()

	if v != -1 {
		t.Errorf("Expected: -1, Got: %v", v)
	}

	if s.ToString() != "[]" {
		t.Errorf("Expected: [], Got: %v", s.ToString())
	}
}

func TestTop(t *testing.T) {
	s := data_structure.NewStack()
	s.Push(5)
	s.Push(3)
	v := s.Top()

	if v != 3 {
		t.Errorf("Expected: 3, Got: %v", v)
	}

	if s.ToString() != "[5 3]" {
		t.Errorf("Expected: [5 3], Got: %v", s.ToString())
	}
}

func TestTopOnEmptyStack(t *testing.T) {
	s := data_structure.NewStack()
	v := s.Top()

	if v != -1 {
		t.Errorf("Expected: -1, Got: %v", v)
	}

	if s.ToString() != "[]" {
		t.Errorf("Expected: [], Got: %v", s.ToString())
	}
}
