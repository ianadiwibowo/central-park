package queuestruct

import (
	"fmt"
	"strings"
)

type QueueStruct struct {
	Storage []interface{}
}

// NewQueue creates a new empty queue
func NewQueueStruct() *QueueStruct {
	return &QueueStruct{}
}

// Enqueue puts the value to the end of the queue
func (q *QueueStruct) Enqueue(value interface{}) {
	q.Storage = append(q.Storage, value)
}

// Dequeue returns the front value and remove it from the queue
func (q *QueueStruct) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	value := q.Storage[0]
	q.Storage[0] = nil
	q.Storage = q.Storage[1:]

	return value
}

// Peek returns the front value without dequeueing it from the queue
func (q *QueueStruct) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.Storage[0]
}

// Clear empties the queue
func (q *QueueStruct) Clear() {
	q.Storage = []interface{}{}
}

// IsEmpty tells whether the queue is empty (true) or not (false)
func (q *QueueStruct) IsEmpty() bool {
	return len(q.Storage) == 0
}

// Print returns the human-readable format of the queue
func (q *QueueStruct) Print() (printout string) {
	for _, v := range q.Storage {
		printout = fmt.Sprintf("%v %v", printout, v)
	}

	return fmt.Sprintf("[%v]", strings.TrimSpace(printout))
}
