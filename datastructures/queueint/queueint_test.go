package queueint_test

import (
	"testing"

	"github.com/ianadiwibowo/central-park/datastructures/queueint"
	"github.com/stretchr/testify/assert"
)

func TestNewQueueInt(t *testing.T) {
	q := queueint.NewQueueInt()

	assert.Equal(t, "[]", q.Print())
}

func TestEnqueue(t *testing.T) {
	q := queueint.NewQueueInt()
	q.Enqueue(5)
	q.Enqueue(3)

	assert.Equal(t, "[5 3]", q.Print())
}

func TestDequeue(t *testing.T) {
	q := queueint.NewQueueInt()
	q.Enqueue(5)
	q.Enqueue(3)
	q.Enqueue(1)
	v := q.Dequeue()

	assert.Equal(t, v, 5)
	assert.Equal(t, "[3 1]", q.Print())
}

func TestDequeueOnEmptyQueue(t *testing.T) {
	q := queueint.NewQueueInt()
	v := q.Dequeue()

	assert.Equal(t, v, -1)
	assert.Equal(t, "[]", q.Print())
}

func TestPeek(t *testing.T) {
	q := queueint.NewQueueInt()
	q.Enqueue(1)
	q.Enqueue(2)
	v := q.Peek()

	assert.Equal(t, v, 1)
	assert.Equal(t, "[1 2]", q.Print())
}

func TestPeekOnEmptyQueue(t *testing.T) {
	q := queueint.NewQueueInt()
	v := q.Peek()

	assert.Equal(t, v, -1)
	assert.Equal(t, "[]", q.Print())
}

func TestClear(t *testing.T) {
	q := queueint.NewQueueInt()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Clear()

	assert.Equal(t, "[]", q.Print())
}

func TestIsEmptyOnEmptyQueue(t *testing.T) {
	q := queueint.NewQueueInt()

	assert.True(t, q.IsEmpty())
}

func TestIsEmptyOnNotEmptyQueue(t *testing.T) {
	q := queueint.NewQueueInt()
	q.Enqueue(1)

	assert.False(t, q.IsEmpty())
}
