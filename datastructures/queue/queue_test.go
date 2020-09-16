package queue_test

import (
	"testing"

	"github.com/ianadiwibowo/central-park/datastructures/queue"
	"github.com/stretchr/testify/assert"
)

// cat demonstrates that a custom struct can be used as the queue internal data structure
type cat struct {
	ID   int
	Name string
}

func TestNewQueue(t *testing.T) {
	q := queue.NewQueue()

	assert.Equal(t, "[]", q.Print())
}

func TestEnqueue(t *testing.T) {
	q := queue.NewQueue()
	q.Enqueue(&cat{1, "Pupuru"})
	q.Enqueue(&cat{3, "Hosico"})

	assert.Equal(t, "[&{1 Pupuru} &{3 Hosico}]", q.Print())
}

func TestDequeue(t *testing.T) {
	q := queue.NewQueue()
	q.Enqueue(&cat{5, "Yuen Jumbo"})
	q.Enqueue(&cat{3, "Fomo"})
	q.Enqueue(&cat{1, "Ah Fei"})
	v := q.Dequeue().(*cat)

	assert.Equal(t, v, &cat{5, "Yuen Jumbo"})
	assert.Equal(t, q.Print(), "[&{3 Fomo} &{1 Ah Fei}]")
}

func TestDequeueOnEmptyQueue(t *testing.T) {
	q := queue.NewQueue()
	v := q.Dequeue()

	assert.Nil(t, v)
	assert.Equal(t, "[]", q.Print())
}

func TestPeek(t *testing.T) {
	q := queue.NewQueue()
	q.Enqueue(&cat{1, "Lupita"})
	q.Enqueue(&cat{2, "Bombi"})
	v := q.Peek().(*cat)

	assert.Equal(t, v, &cat{1, "Lupita"})
	assert.Equal(t, "[&{1 Lupita} &{2 Bombi}]", q.Print())
}

func TestPeekOnEmptyQueue(t *testing.T) {
	q := queue.NewQueue()
	v := q.Peek()

	assert.Nil(t, v)
	assert.Equal(t, "[]", q.Print())
}

func TestClear(t *testing.T) {
	q := queue.NewQueue()
	q.Enqueue(&cat{10, "Simabossneko"})
	q.Enqueue(&cat{20, "Big Billy"})
	q.Clear()

	assert.Equal(t, "[]", q.Print())
}

func TestIsEmptyOnEmptyQueue(t *testing.T) {
	q := queue.NewQueue()

	assert.True(t, q.IsEmpty())
}

func TestIsEmptyOnNotEmptyQueue(t *testing.T) {
	q := queue.NewQueue()
	q.Enqueue(&cat{7, "Bellamy"})

	assert.False(t, q.IsEmpty())
}
