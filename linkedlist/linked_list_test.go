package linkedlist_test

import (
	"testing"

	linkedlist "github.com/ianadiwibowo/central-park/linkedlist"
)

func TestLinkedListNewLinkedList(t *testing.T) {
	l := linkedlist.NewLinkedList()

	if l.Print() != "[]" {
		t.Errorf("Expected: []. Got: %v", l.Print())
	}
}

func TestLinkedListAdd(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(7)
	l.Add(10)
	l.Add(30)

	if l.Print() != "[7 10 30]" {
		t.Errorf("Expected: [7 10 30]. Got: %v", l.Print())
	}
}

func TestLinkedListInsertBefore(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertBefore(3, 2)
	if l.Print() != "[1 3 2]" {
		t.Errorf("Expected: [1 3 2]. Got: %v", l.Print())
	}
}

func TestLinkedListInsertBeforeNonExistentValue(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertBefore(3, 5)
	if l.Print() != "[1 2]" {
		t.Errorf("Expected: [1 2]. Got: %v", l.Print())
	}
}

func TestLinkedListInsertBeforeFirstNode(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertBefore(10, 1)
	if l.Print() != "[10 1 2]" {
		t.Errorf("Expected: [10 1 2]. Got: %v", l.Print())
	}
}

func TestLinkedListInsertAfter(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertAfter(3, 1)
	if l.Print() != "[1 3 2]" {
		t.Errorf("Expected: [1 3 2]. Got: %v", l.Print())
	}
}

func TestLinkedListInsertAfterNonExistentValue(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertAfter(3, 5)
	if l.Print() != "[1 2]" {
		t.Errorf("Expected: [1 2]. Got: %v", l.Print())
	}
}

func TestLinkedListInsertAfterLastNode(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(1)
	l.Add(2)

	l.InsertAfter(10, 2)
	if l.Print() != "[1 2 10]" {
		t.Errorf("Expected: [1 2 10]. Got: %v", l.Print())
	}
}

func TestLinkedListDelete(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(8)

	l.Delete(6)
	if l.Print() != "[2 4 8]" {
		t.Errorf("Expected: [2 4 8]. Got: %v", l.Print())
	}
}

func TestLinkedListDeleteFirstNode(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(8)

	l.Delete(2)
	if l.Print() != "[4 6 8]" {
		t.Errorf("Expected: [4 6 8]. Got: %v", l.Print())
	}
}

func TestDeleteLastNode(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(8)

	l.Delete(8)
	if l.Print() != "[2 4 6]" {
		t.Errorf("Expected: [2 4 6]. Got: %v", l.Print())
	}
}

func TestDeleteWithNonExistentValue(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(8)

	l.Delete(5)
	if l.Print() != "[2 4 6 8]" {
		t.Errorf("Expected: [2 4 6 8]. Got: %v", l.Print())
	}
}

func TestLinkedListFind(t *testing.T) {
	l := linkedlist.NewLinkedList()
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
	l := linkedlist.NewLinkedList()
	l.Add(100)
	l.Add(200)
	l.Add(300)

	node := l.Find(400)
	if node != nil {
		t.Errorf("Expected: nil. Got: %v", node.Value)
	}
}
