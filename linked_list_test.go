package data_structure_test

import (
	"testing"

	data_structure "github.com/ianadiwibowo/central-park"
)

func TestNewLinkedList(t *testing.T) {
	l := data_structure.NewLinkedList()

	if l.ToString() != "[]" {
		t.Errorf("Expected: []. Got: %v", l.ToString())
	}
}

func TestAdd(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(7)
	l.Add(10)
	l.Add(30)

	if l.ToString() != "[7 10 30]" {
		t.Errorf("Expected: [7 10 30]. Got: %v", l.ToString())
	}
}
