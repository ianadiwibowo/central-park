package data_structure_test

import (
	"testing"

	data_structure "github.com/ianadiwibowo/central-park"
)

func TestQueueNewQueue(t *testing.T) {
	q := data_structure.NewQueue()

	if q.Print() != "[]" {
		t.Errorf("Expected: []. Got: %v", q.Print())
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := data_structure.NewQueue()
	q.Enqueue(5)
	q.Enqueue(3)

	if q.Print() != "[5 3]" {
		t.Errorf("Expected: [5 3], Got: %v", q.Print())
	}
}

func TestQueueDequeue(t *testing.T) {
	q := data_structure.NewQueue()
	q.Enqueue(5)
	q.Enqueue(3)
	q.Enqueue(1)
	v := q.Dequeue()

	if v != 5 {
		t.Errorf("Expected: 5, Got: %v", q.Print())
	}

	if q.Print() != "[3 1]" {
		t.Errorf("Expected: [3 1], Got: %v", q.Print())
	}
}

func TestQueueDequeueOnEmptyQueue(t *testing.T) {
	q := data_structure.NewQueue()
	v := q.Dequeue()

	if v != -1 {
		t.Errorf("Expected: -1, Got: %v", q.Print())
	}

	if q.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", q.Print())
	}
}

func TestQueuePeek(t *testing.T) {
	q := data_structure.NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	v := q.Peek()

	if v != 1 {
		t.Errorf("Expected: 1, Got: %v", q.Print())
	}

	if q.Print() != "[1 2]" {
		t.Errorf("Expected: [1 2], Got: %v", q.Print())
	}
}

func TestQueuePeekOnEmptyQueue(t *testing.T) {
	q := data_structure.NewQueue()
	v := q.Peek()

	if v != -1 {
		t.Errorf("Expected: -1, Got: %v", q.Print())
	}

	if q.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", q.Print())
	}
}

func TestQueueClear(t *testing.T) {
	q := data_structure.NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Clear()

	if q.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", q.Print())
	}
}

func TestQueueIsEmptyOnEmptyQueue(t *testing.T) {
	q := data_structure.NewQueue()

	if q.IsEmpty() == false {
		t.Errorf("Expected: true, Got: %v", q.Print())
	}
}

func TestQueueIsEmptyOnNotEmptyQueue(t *testing.T) {
	q := data_structure.NewQueue()
	q.Enqueue(1)

	if q.IsEmpty() == true {
		t.Errorf("Expected: false, Got: %v", q.Print())
	}
}
