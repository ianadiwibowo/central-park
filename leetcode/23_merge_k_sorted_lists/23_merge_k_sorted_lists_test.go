package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	a.Equal(t, true, true)
}

func TestCreateLinkedList(t *testing.T) {
	node1 := createLinkedList([]int{})
	a.Equal(t, (*ListNode)(nil), node1)

	node3 := createLinkedList([]int{5})
	a.Equal(t, 5, node3.Val)
	a.Equal(t, (*ListNode)(nil), node3.Next)

	node2 := createLinkedList([]int{1, 2, 3})
	a.Equal(t, 1, node2.Val)
	a.Equal(t, 2, node2.Next.Val)
	a.Equal(t, 3, node2.Next.Next.Val)
	a.Equal(t, (*ListNode)(nil), node2.Next.Next.Next)
}

func TestPrintLinkedList(t *testing.T) {
	node1 := createLinkedList([]int{})
	a.Equal(t, "", printLinkedList(node1))

	node3 := createLinkedList([]int{5})
	a.Equal(t, "5", printLinkedList(node3))

	node2 := createLinkedList([]int{1, 2, 3})
	a.Equal(t, "1 2 3", printLinkedList(node2))
}
