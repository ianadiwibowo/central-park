package stack_test

import (
	"testing"

	"github.com/ianadiwibowo/central-park/datastructures/stack"
	"github.com/stretchr/testify/assert"
)

func TestStackNewStack(t *testing.T) {
	s := stack.NewStack()

	assert.Equal(t, "[]", s.Print())
}

func TestStackPush(t *testing.T) {
	s := stack.NewStack()
	s.Push(5)
	s.Push(3)

	assert.Equal(t, "[5 3]", s.Print())
}

func TestStackPop(t *testing.T) {
	s := stack.NewStack()
	s.Push(5)
	s.Push(3)
	v := s.Pop()

	assert.Equal(t, v, 3)
	assert.Equal(t, "[5]", s.Print())
}

func TestStackPopOnEmptyStack(t *testing.T) {
	s := stack.NewStack()
	v := s.Pop()

	assert.Equal(t, v, -1)
	assert.Equal(t, "[]", s.Print())
}

func TestStackPeek(t *testing.T) {
	s := stack.NewStack()
	s.Push(5)
	s.Push(3)
	v := s.Peek()

	assert.Equal(t, v, 3)
	assert.Equal(t, "[5 3]", s.Print())
}

func TestStackPeekOnEmptyStack(t *testing.T) {
	s := stack.NewStack()
	v := s.Peek()

	assert.Equal(t, v, -1)
	assert.Equal(t, "[]", s.Print())
}

func TestStackIsEmptyOnEmptyStack(t *testing.T) {
	s := stack.NewStack()

	assert.True(t, s.IsEmpty())
}

func TestStackIsEmptyOnNotEmptyStack(t *testing.T) {
	s := stack.NewStack()
	s.Push(5)
	s.Push(3)

	assert.False(t, s.IsEmpty())
}
