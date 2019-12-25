package queue

import (
	"fmt"
)

type Queue struct {
	Storage []int
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	return &Queue{
		Storage: []int{},
	}
}

// Enqueue puts the value to the end of the queue
func (q *Queue) Enqueue(value int) {
	q.Storage = append(q.Storage, value)
}

// Dequeue returns the front value and remove it from the queue
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}

	value := q.Storage[0]
	q.Storage[0] = 0
	q.Storage = q.Storage[1:]

	return value
}

// Peek returns the front value without dequeueing it from the queue
func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}

	return q.Storage[0]
}

// Clear empties the queue
func (q *Queue) Clear() {
	q.Storage = []int{}
}

// IsEmpty tells whether the queue is empty (true) or not (false)
func (q *Queue) IsEmpty() bool {
	return len(q.Storage) == 0
}

// Print returns the human-readable format of the queue
func (q *Queue) Print() string {
	return fmt.Sprintf("%v", q.Storage)
}
