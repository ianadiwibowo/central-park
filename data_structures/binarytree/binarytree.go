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
	return &BinaryTree{}
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

// FindCompletePaths find all complete paths from root to all leaves
func (b *BinaryTree) FindCompletePaths() [][]int {
	return findCompletePaths(b.Root, []int{}, [][]int{})
}

func findCompletePaths(currentNode *BinaryTreeNode, path []int, prevResults [][]int) [][]int {
	if currentNode != nil {
		updatedPath := make([]int, len(path)+1)
		copy(updatedPath, append(path, currentNode.Value))

		if currentNode.IsLeaf() {
			updatedPrevResults := make([][]int, len(prevResults)+1)
			copy(updatedPrevResults, append(prevResults, updatedPath))

			return updatedPrevResults
		}
		prevResults = findCompletePaths(currentNode.Left, updatedPath, prevResults)
		prevResults = findCompletePaths(currentNode.Right, updatedPath, prevResults)
	}

	return prevResults
}

// PrintPreOrder returns the pre-ordered (depth first) human-readable format of the binary tree
func (b *BinaryTree) PrintPreOrder() string {
	return fmt.Sprintf("[%v]", strings.TrimSpace(printPreOrder(b.Root, "")))
}

func printPreOrder(currentNode *BinaryTreeNode, prevPrintout string) (printout string) {
	if currentNode != nil {
		printout = fmt.Sprintf("%v %v", prevPrintout, currentNode.Value)
		printout := printPreOrder(currentNode.Left, printout)
		return printPreOrder(currentNode.Right, printout)
	}

	return prevPrintout
}

// PrintInOrder returns the in-ordered (depth first) human-readable format of the binary tree
func (b *BinaryTree) PrintInOrder() string {
	return fmt.Sprintf("[%v]", strings.TrimSpace(printInOrder(b.Root, "")))
}

func printInOrder(currentNode *BinaryTreeNode, prevPrintout string) string {
	if currentNode != nil {
		printout := printInOrder(currentNode.Left, prevPrintout)
		printout = fmt.Sprintf("%v %v", printout, currentNode.Value)
		return printInOrder(currentNode.Right, printout)
	}

	return prevPrintout
}

// PrintPostOrder returns the post-ordered (depth first) human-readable format of the binary tree
func (b *BinaryTree) PrintPostOrder() string {
	return fmt.Sprintf("[%v]", strings.TrimSpace(printPostOrder(b.Root, "")))
}

func printPostOrder(currentNode *BinaryTreeNode, prevPrintout string) string {
	if currentNode != nil {
		printout := printPostOrder(currentNode.Left, prevPrintout)
		printout = printPostOrder(currentNode.Right, printout)
		return fmt.Sprintf("%v %v", printout, currentNode.Value)
	}

	return prevPrintout
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
func (btn *BinaryTreeNode) IsLeaf() bool {
	return btn.Left == nil && btn.Right == nil
}
