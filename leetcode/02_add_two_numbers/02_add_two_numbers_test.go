package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := printLinkedList(addTwoNumbers(l1, l2))
	a.Equal(t, "708", result)

	l1 = &ListNode{1, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}
	l2 = &ListNode{1, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}
	result = printLinkedList(addTwoNumbers(l1, l2))
	a.Equal(t, "2002", result)

	l1 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	result = printLinkedList(addTwoNumbers(l1, l2))
	a.Equal(t, "89990001", result)

	l1 = &ListNode{0, nil}
	l2 = &ListNode{0, nil}
	result = printLinkedList(addTwoNumbers(l1, l2))
	a.Equal(t, "0", result)

	l1 = &ListNode{5, nil}
	l2 = &ListNode{5, nil}
	result = printLinkedList(addTwoNumbers(l1, l2))
	a.Equal(t, "01", result)
}
