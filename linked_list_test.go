package data_structure_test

import (
	"testing"

	data_structure "github.com/ianadiwibowo/central-park"
)

func TestLinkedListNewLinkedList(t *testing.T) {
	l := data_structure.NewLinkedList()

	if l.ToString() != "[]" {
		t.Errorf("Expected: []. Got: %v", l.ToString())
	}
}

func TestLinkedListAdd(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(7)
	l.Add(10)
	l.Add(30)

	if l.ToString() != "[7 10 30]" {
		t.Errorf("Expected: [7 10 30]. Got: %v", l.ToString())
	}
}

func TestLinkedListInsertBefore(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertBefore(3, 2)
	if l.ToString() != "[1 3 2]" {
		t.Errorf("Expected: [1 3 2]. Got: %v", l.ToString())
	}
}

func TestLinkedListInsertBeforeNonExistentValue(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertBefore(3, 5)
	if l.ToString() != "[1 2]" {
		t.Errorf("Expected: [1 2]. Got: %v", l.ToString())
	}
}

func TestLinkedListInsertBeforeFirstNode(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertBefore(10, 1)
	if l.ToString() != "[10 1 2]" {
		t.Errorf("Expected: [10 1 2]. Got: %v", l.ToString())
	}
}

func TestLinkedListInsertAfter(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertAfter(3, 1)
	if l.ToString() != "[1 3 2]" {
		t.Errorf("Expected: [1 3 2]. Got: %v", l.ToString())
	}
}

func TestLinkedListInsertAfterNonExistentValue(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertAfter(3, 5)
	if l.ToString() != "[1 2]" {
		t.Errorf("Expected: [1 2]. Got: %v", l.ToString())
	}
}

func TestLinkedListInsertAfterLastNode(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertAfter(10, 2)
	if l.ToString() != "[1 2 10]" {
		t.Errorf("Expected: [1 2 10]. Got: %v", l.ToString())
	}
}

func TestLinkedListDelete(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(8)

	l.Delete(6)
	if l.ToString() != "[2 4 8]" {
		t.Errorf("Expected: [2 4 8]. Got: %v", l.ToString())
	}
}

func TestLinkedListDeleteFirstNode(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(8)

	l.Delete(2)
	if l.ToString() != "[4 6 8]" {
		t.Errorf("Expected: [4 6 8]. Got: %v", l.ToString())
	}
}

func TestDeleteLastNode(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(8)

	l.Delete(8)
	if l.ToString() != "[2 4 6]" {
		t.Errorf("Expected: [2 4 6]. Got: %v", l.ToString())
	}
}

func TestDeleteWithNonExistentValue(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(8)

	l.Delete(5)
	if l.ToString() != "[2 4 6 8]" {
		t.Errorf("Expected: [2 4 6 8]. Got: %v", l.ToString())
	}
}

func TestLinkedListFind(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(100)
	l.Add(200)
	l.Add(300)

	node := l.Find(300)
	if node.Value != 300 {
		t.Errorf("Expected: 300. Got: %v", node.Value)
	}

	node = l.Find(100)
	if node.Value != 100 {
		t.Errorf("Expected: 100. Got: %v", node.Value)
	}
}

func TestLinkedListFindWithNonExistentValue(t *testing.T) {
	l := data_structure.NewLinkedList()
	l.Add(100)
	l.Add(200)
	l.Add(300)

	node := l.Find(400)
	if node != nil {
		t.Errorf("Expected: nil. Got: %v", node.Value)
	}
}
