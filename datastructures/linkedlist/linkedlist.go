package linkedlist

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	Head *LinkedListNode
}

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Add puts the value to a new node at the end of the linked list
func (l *LinkedList) Add(value int) {
	newNode := &LinkedListNode{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
}

// InsertBefore puts the value before nextValue's node of the linked list
func (l *LinkedList) InsertBefore(value int, nextValue int) {
	previousNode := &LinkedListNode{}
	nextNode := l.Head

	for nextNode != nil {
		if nextNode.Value == nextValue {
			newNode := &LinkedListNode{
				Value: value,
				Next:  nextNode,
			}
			if previousNode.Next == nil {
				l.Head = newNode
			} else {
				previousNode.Next = newNode
			}
			return
		} else {
			previousNode = nextNode
			nextNode = nextNode.Next
		}
	}
}

// InsertAfter puts the value after previousValue's node of the linked list
func (l *LinkedList) InsertAfter(value int, previousValue int) {
	previousNode := l.Find(previousValue)
	if previousNode != nil {
		newNode := &LinkedListNode{
			Value: value,
			Next:  previousNode.Next,
		}
		previousNode.Next = newNode
	}
}

// Delete removes the value's node from the linked list
func (l *LinkedList) Delete(value int) {
	previousNode := &LinkedListNode{}
	currentNode := l.Head
	nextNode := l.Head.Next

	for currentNode != nil {
		if currentNode.Value == value {
			if previousNode.Next == nil {
				l.Head = nextNode
			} else {
				previousNode.Next = nextNode
			}
			currentNode = nil
			return
		} else {
			previousNode = currentNode
			currentNode = nextNode
			if nextNode != nil {
				nextNode = nextNode.Next
			}
		}
	}
}

// Find returns the value's node in the linked list
func (l *LinkedList) Find(value int) (node *LinkedListNode) {
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}

// Print returns the human-readable format of the linked list
func (l *LinkedList) Print() string {
	v := ""
	currentNode := l.Head
	for currentNode != nil {
		v = fmt.Sprintf("%v%v ", v, currentNode.Value)
		currentNode = currentNode.Next
	}
	v = fmt.Sprintf("[%v]", strings.TrimSpace(v))

	return v
}

