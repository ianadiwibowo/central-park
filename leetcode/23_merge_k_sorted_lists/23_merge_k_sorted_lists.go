package main

// 23. Merge k Sorted Lists (Hard)
// https://leetcode.com/problems/merge-k-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return nil
}

func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	root := &ListNode{}
	currentNode := root
	for i, n := range values {
		currentNode.Val = n // 1
		if i < len(values)-1 {
			currentNode.Next = &ListNode{}
			currentNode = currentNode.Next
		}
	}

	return root
}
