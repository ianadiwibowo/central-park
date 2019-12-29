package queueint

import (
	"fmt"
)

type QueueInt struct {
	Storage []int
}

// NewQueueInt creates a new empty queue
func NewQueueInt() *QueueInt {
	return &QueueInt{}
}

// Enqueue puts the value to the end of the queue
func (q *QueueInt) Enqueue(value int) {
	q.Storage = append(q.Storage, value)
}

// Dequeue returns the front value and remove it from the queue
func (q *QueueInt) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}

	value := q.Storage[0]
	q.Storage[0] = 0
	q.Storage = q.Storage[1:]

	return value
}

// Peek returns the front value without dequeueing it from the queue
func (q *QueueInt) Peek() int {
	if q.IsEmpty() {
		return -1
	}

	return q.Storage[0]
}

// Clear empties the queue
func (q *QueueInt) Clear() {
	q.Storage = []int{}
}

// IsEmpty tells whether the queue is empty (true) or not (false)
func (q *QueueInt) IsEmpty() bool {
	return len(q.Storage) == 0
}

// Print returns the human-readable format of the queue
func (q *QueueInt) Print() string {
	return fmt.Sprintf("%v", q.Storage)
}
