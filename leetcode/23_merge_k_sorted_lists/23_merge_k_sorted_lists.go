package main

import "fmt"

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

func printLinkedList(root *ListNode) string {
	if root == nil {
		return ""
	}

	node := root
	result := ""
	for {
		result = fmt.Sprintf("%s%v", result, node.Val)
		if node.Next == nil {
			break
		}
		result = result + " "
		node = node.Next
	}
	return result
}
