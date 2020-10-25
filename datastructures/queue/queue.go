package queue

import (
	"fmt"
	"strings"
)

type Queue struct {
	Storage []interface{}
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue puts the value to the end of the queue
func (q *Queue) Enqueue(value interface{}) {
	q.Storage = append(q.Storage, value)
}

// Dequeue returns the front value and remove it from the queue
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	value := q.Storage[0]
	q.Storage = q.Storage[1:]

	return value
}

// Peek returns the front value without dequeueing it from the queue
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.Storage[0]
}

// Clear empties the queue
func (q *Queue) Clear() {
	q.Storage = []interface{}{}
}

// IsEmpty tells whether the queue is empty (true) or not (false)
func (q *Queue) IsEmpty() bool {
	return len(q.Storage) == 0
}

// Print returns the human-readable format of the queue
func (q *Queue) Print() (printout string) {
	for _, v := range q.Storage {
		printout = fmt.Sprintf("%v %v", printout, v)
	}

	return fmt.Sprintf("[%v]", strings.TrimSpace(printout))
}
