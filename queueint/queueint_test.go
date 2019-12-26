package queueint_test

import (
	"testing"

	"github.com/ianadiwibowo/central-park/queueint"
)

func TestNewQueueInt(t *testing.T) {
	q := queueint.NewQueueInt()

	if q.Print() != "[]" {
		t.Errorf("Expected: []. Got: %v", q.Print())
	}
}

func TestEnqueue(t *testing.T) {
	q := queueint.NewQueueInt()
	q.Enqueue(5)
	q.Enqueue(3)

	if q.Print() != "[5 3]" {
		t.Errorf("Expected: [5 3], Got: %v", q.Print())
	}
}

func TestDequeue(t *testing.T) {
	q := queueint.NewQueueInt()
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

func TestDequeueOnEmptyQueue(t *testing.T) {
	q := queueint.NewQueueInt()
	v := q.Dequeue()

	if v != -1 {
		t.Errorf("Expected: -1, Got: %v", q.Print())
	}

	if q.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", q.Print())
	}
}

func TestPeek(t *testing.T) {
	q := queueint.NewQueueInt()
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

func TestPeekOnEmptyQueue(t *testing.T) {
	q := queueint.NewQueueInt()
	v := q.Peek()

	if v != -1 {
		t.Errorf("Expected: -1, Got: %v", q.Print())
	}

	if q.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", q.Print())
	}
}

func TestClear(t *testing.T) {
	q := queueint.NewQueueInt()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Clear()

	if q.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", q.Print())
	}
}

func TestIsEmptyOnEmptyQueue(t *testing.T) {
	q := queueint.NewQueueInt()

	if q.IsEmpty() == false {
		t.Errorf("Expected: true, Got: %v", q.Print())
	}
}

func TestIsEmptyOnNotEmptyQueue(t *testing.T) {
	q := queueint.NewQueueInt()
	q.Enqueue(1)

	if q.IsEmpty() == true {
		t.Errorf("Expected: false, Got: %v", q.Print())
	}
}
