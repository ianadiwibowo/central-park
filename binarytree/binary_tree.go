package binarytree

import (
	"fmt"
	"strings"

	"github.com/ianadiwibowo/central-park/queue"
)

type BinaryTree struct {
	Root *BinaryTreeNode
}

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// NewBinaryTree creates a new empty binary tree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		Root: nil,
	}
}

// SetRoot assign a new value's node as the tree root
func (b *BinaryTree) SetRoot(value int) {
	b.Root = &BinaryTreeNode{
		Value: value,
	}
}

// InsertLeft create a new value's node and put it as left child of parent's node
func (b *BinaryTree) InsertLeft(value int, parent int) {
	parentNode := b.Find(parent)
	if parentNode != nil {
		parentNode.Left = &BinaryTreeNode{
			Value: value,
		}
	}
}

// InsertRight create a new value's node and put it as right child of parent's node
func (b *BinaryTree) InsertRight(value int, parent int) {
	parentNode := b.Find(parent)
	if parentNode != nil {
		parentNode.Right = &BinaryTreeNode{
			Value: value,
		}
	}
}

// Find returns value's node
func (b *BinaryTree) Find(value int) *BinaryTreeNode {
	return findInOrder(b.Root, value)
}

func findInOrder(currentNode *BinaryTreeNode, value int) *BinaryTreeNode {
	if currentNode != nil {
		leftNode := findInOrder(currentNode.Left, value)
		if leftNode != nil {
			return leftNode
		}

		if currentNode.Value == value {
			return currentNode
		}

		rightNode := findInOrder(currentNode.Right, value)
		if rightNode != nil {
			return rightNode
		}
	}

	return nil
}

// Height returns the tree height from the root
func (b *BinaryTree) Height() int {
	return maxHeightPreOrder(b.Root, 0)
}

func maxHeightPreOrder(currentNode *BinaryTreeNode, maxHeight int) int {
	if currentNode != nil {
		leftHeight := maxHeightPreOrder(currentNode.Left, maxHeight+1)
		rightHeight := maxHeightPreOrder(currentNode.Right, maxHeight+1)
		if leftHeight > maxHeight {
			maxHeight = leftHeight
		}
		if rightHeight > maxHeight {
			maxHeight = rightHeight
		}
		return maxHeight
	}

	return 0
}

// PrintPreOrder returns the pre-ordered (depth first) human-readable format of the binary tree
func (b *BinaryTree) PrintPreOrder() string {
	return fmt.Sprintf("[%v]", strings.TrimSpace(printPreOrder(b.Root, "")))
}

func printPreOrder(currentNode *BinaryTreeNode, previousPrintout string) (printout string) {
	if currentNode != nil {
		printout = fmt.Sprintf("%v %v", previousPrintout, currentNode.Value)
		leftPrintout := printPreOrder(currentNode.Left, printout)
		return printPreOrder(currentNode.Right, leftPrintout)
	}

	return previousPrintout
}

// PrintInOrder returns the in-ordered (depth first) human-readable format of the binary tree
func (b *BinaryTree) PrintInOrder() string {
	return fmt.Sprintf("[%v]", strings.TrimSpace(printInOrder(b.Root, "")))
}

func printInOrder(currentNode *BinaryTreeNode, previousPrintout string) string {
	if currentNode != nil {
		leftPrintout := printInOrder(currentNode.Left, previousPrintout)
		currentPrintout := fmt.Sprintf("%v %v", leftPrintout, currentNode.Value)
		return printInOrder(currentNode.Right, currentPrintout)
	}

	return previousPrintout
}

// PrintPostOrder returns the post-ordered (depth first) human-readable format of the binary tree
func (b *BinaryTree) PrintPostOrder() string {
	return fmt.Sprintf("[%v]", strings.TrimSpace(printPostOrder(b.Root, "")))
}

func printPostOrder(currentNode *BinaryTreeNode, previousPrintout string) string {
	if currentNode != nil {
		leftPrintout := printPostOrder(currentNode.Left, previousPrintout)
		rightPrintout := printPostOrder(currentNode.Right, leftPrintout)
		return fmt.Sprintf("%v %v", rightPrintout, currentNode.Value)
	}

	return previousPrintout
}

// PrintLevelOrder returns the level-ordered (breadth first) human-readable format of the binary tree
func (b *BinaryTree) PrintLevelOrder() (printout string) {
	queue := queue.NewQueue()
	queue.Enqueue(b.Root)

	for !queue.IsEmpty() {
		currentNode := queue.Dequeue().(*BinaryTreeNode)
		printout = fmt.Sprintf("%v %v", printout, currentNode.Value)

		if currentNode.Left != nil {
			queue.Enqueue(currentNode.Left)
		}

		if currentNode.Right != nil {
			queue.Enqueue(currentNode.Right)
		}
	}

	return fmt.Sprintf("[%v]", strings.TrimSpace(printout))
}

// IsLeaf checks whether a binary tree node is leaf (true) or not (false)
func (b *BinaryTreeNode) IsLeaf() bool {
	return b.Left == nil && b.Right == nil
}
