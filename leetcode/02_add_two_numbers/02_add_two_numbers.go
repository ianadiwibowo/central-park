package main

import "fmt"

// 2. Add Two Numbers (Medium)
// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	r := result
	remainder := 0
	l1Ended := false
	l2Ended := false

	for {
		a := 0
		if !l1Ended {
			a = l1.Val
		}

		b := 0
		if !l2Ended {
			b = l2.Val
		}

		sum := a + b + remainder
		if sum > 9 {
			sum = sum % 10
			remainder = 1
		} else {
			remainder = 0
		}
		r.Val = sum

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1Ended = true
		}

		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2Ended = true
		}

		if l1Ended && l2Ended && remainder == 0 {
			break
		}

		r.Next = &ListNode{}
		r = r.Next
	}

	return result
}

func printLinkedList(l *ListNode) string {
	result := fmt.Sprintf("%v", l.Val)
	for l.Next != nil {
		l = l.Next
		result = fmt.Sprintf("%v%v", result, l.Val)
	}
	return result
}
