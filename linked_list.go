package data_structure

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	head *LinkedListNode
}

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Add puts the value at the end of the linked list
func (l *LinkedList) Add(value int) {
	newNode := &LinkedListNode{
		value: value,
		next:  nil,
	}

	if l.head == nil {
		l.head = newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

// Insert puts the value at the position index of the linked list
func (l *LinkedList) Insert(value int, position int) {

}

// Delete removes the value from the linked list
func (l *LinkedList) Delete(value int) {

}

// Find returns the position index of the value in the linked list
func (l *LinkedList) Find(value int) (position int) {
	return 0
}

// ToString returns the human-readable format of the stack
func (l *LinkedList) ToString() string {
	v := ""
	currentNode := l.head
	for currentNode != nil {
		v = fmt.Sprintf("%v%v ", v, currentNode.value)
		currentNode = currentNode.next
	}
	v = fmt.Sprintf("[%v]", strings.TrimSpace(v))

	return v
}
