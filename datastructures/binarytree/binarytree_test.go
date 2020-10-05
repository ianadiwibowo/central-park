package binarytree_test

import (
	"testing"

	"github.com/ianadiwibowo/central-park/datastructures/binarytree"
	"github.com/stretchr/testify/assert"
)

func TestNewBinaryTree(t *testing.T) {
	b := binarytree.NewBinaryTree()

	assert.Equal(t, []int(nil), b.TraversePreOrder())
}

func TestSetRoot(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(5)

	assert.Equal(t, []int{5}, b.TraversePreOrder())
}

func TestInsertLeft(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(5)
	b.InsertLeft(1, 5)

	assert.Equal(t, []int{5, 1}, b.TraversePreOrder())
}

func TestInsertLeftOnNonExistentValue(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(5)
	b.InsertLeft(1, 4)

	assert.Equal(t, []int{5}, b.TraversePreOrder())
}

func TestInsertRight(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(5)
	b.InsertRight(7, 5)

	assert.Equal(t, []int{5, 7}, b.TraversePreOrder())
}

func TestInsertRightOnNonExistentValue(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(5)
	b.InsertRight(7, 4)

	assert.Equal(t, []int{5}, b.TraversePreOrder())
}

func TestFind(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(5)
	b.InsertRight(7, 5)
	currentNode := b.Find(7)

	assert.Equal(t, 7, currentNode.Value)
}

func TestFindNonExistentValue(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(5)
	b.InsertRight(7, 5)
	currentNode := b.Find(8)

	assert.Nil(t, currentNode)
}

func TestHeightOnBalancedTree(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(5)
	b.InsertLeft(1, 5)
	b.InsertRight(2, 5)

	assert.Equal(t, 1, b.Height())
}

func TestHeightOnBiggerBalancedTree(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(1)
	b.InsertLeft(2, 1)
	b.InsertLeft(3, 2)
	b.InsertRight(4, 2)
	b.InsertRight(5, 1)
	b.InsertLeft(6, 5)
	b.InsertRight(7, 5)

	assert.Equal(t, 2, b.Height())
}

func TestHeightOnSkewedLeftTree(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(1)
	b.InsertLeft(2, 1)
	b.InsertLeft(3, 2)
	b.InsertLeft(4, 3)
	b.InsertLeft(5, 4)

	assert.Equal(t, 4, b.Height())
}

func TestHeightOnSkewedRightTree(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(1)
	b.InsertRight(2, 1)
	b.InsertRight(3, 2)
	b.InsertRight(4, 3)
	b.InsertRight(5, 4)

	assert.Equal(t, 4, b.Height())
}

func TestHeightOnRootOnlyTree(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(1)

	assert.Equal(t, 0, b.Height())
}

func TestHeightOnEmptyTree(t *testing.T) {
	b := binarytree.NewBinaryTree()

	assert.Equal(t, 0, b.Height())
}

func TestFindCompletePaths(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(1)
	b.InsertLeft(2, 1)
	b.InsertLeft(4, 2)
	b.InsertRight(5, 2)
	b.InsertRight(3, 1)
	b.InsertLeft(6, 3)
	b.InsertRight(7, 3)
	expected := [][]int{
		{1, 2, 4},
		{1, 2, 5},
		{1, 3, 6},
		{1, 3, 7},
	}

	result := b.FindCompletePaths()

	assert.Equal(t, expected, result)
}

func TestFindCompletePathsWithHalfLeaves(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(1)
	b.InsertLeft(2, 1)
	b.InsertLeft(3, 2)
	b.InsertLeft(4, 3)
	b.InsertRight(5, 2)
	b.InsertLeft(6, 5)
	b.InsertRight(7, 1)
	b.InsertLeft(8, 7)
	b.InsertLeft(9, 8)
	b.InsertRight(10, 7)
	b.InsertLeft(11, 10)
	expected := [][]int{
		{1, 2, 3, 4},
		{1, 2, 5, 6},
		{1, 7, 8, 9},
		{1, 7, 10, 11},
	}
	result := b.FindCompletePaths()

	assert.Equal(t, expected, result)
}

func TestFindCompletePathsWithBiggerCompleteLeaves(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(1)
	b.InsertLeft(2, 1)
	b.InsertLeft(3, 2)
	b.InsertLeft(4, 3)
	b.InsertRight(12, 3)
	b.InsertRight(5, 2)
	b.InsertLeft(6, 5)
	b.InsertRight(13, 5)
	b.InsertRight(7, 1)
	b.InsertLeft(8, 7)
	b.InsertLeft(9, 8)
	b.InsertRight(14, 8)
	b.InsertRight(10, 7)
	b.InsertLeft(11, 10)
	b.InsertRight(15, 10)
	expected := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 12},
		{1, 2, 5, 6},
		{1, 2, 5, 13},
		{1, 7, 8, 9},
		{1, 7, 8, 14},
		{1, 7, 10, 11},
		{1, 7, 10, 15},
	}

	result := b.FindCompletePaths()

	assert.Equal(t, expected, result)
}

func TestTraversePreOrder(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(8)
	b.InsertLeft(5, 8)
	b.InsertLeft(9, 5)
	b.InsertRight(7, 5)
	b.InsertLeft(1, 7)
	b.InsertRight(12, 7)
	b.InsertLeft(2, 12)
	b.InsertRight(4, 8)
	b.InsertRight(11, 4)
	b.InsertLeft(3, 11)

	assert.Equal(t, []int{8, 5, 9, 7, 1, 12, 2, 4, 11, 3}, b.TraversePreOrder())
}

func TestTraverseInOrder(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(8)
	b.InsertLeft(5, 8)
	b.InsertLeft(9, 5)
	b.InsertRight(7, 5)
	b.InsertLeft(1, 7)
	b.InsertRight(12, 7)
	b.InsertLeft(2, 12)
	b.InsertRight(4, 8)
	b.InsertRight(11, 4)
	b.InsertLeft(3, 11)

	assert.Equal(t, []int{9, 5, 1, 7, 2, 12, 8, 4, 3, 11}, b.TraverseInOrder())
}

func TestTraversePostOrder(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(8)
	b.InsertLeft(5, 8)
	b.InsertLeft(9, 5)
	b.InsertRight(7, 5)
	b.InsertLeft(1, 7)
	b.InsertRight(12, 7)
	b.InsertLeft(2, 12)
	b.InsertRight(4, 8)
	b.InsertRight(11, 4)
	b.InsertLeft(3, 11)

	assert.Equal(t, []int{9, 1, 2, 12, 7, 5, 3, 11, 4, 8}, b.TraversePostOrder())
}

func TestTraverseLevelOrder(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(8)
	b.InsertLeft(5, 8)
	b.InsertLeft(9, 5)
	b.InsertRight(7, 5)
	b.InsertLeft(1, 7)
	b.InsertRight(12, 7)
	b.InsertLeft(2, 12)
	b.InsertRight(4, 8)
	b.InsertRight(11, 4)
	b.InsertLeft(3, 11)

	assert.Equal(t, []int{8, 5, 4, 9, 7, 11, 1, 12, 3, 2}, b.TraverseLevelOrder())
}

func TestIsLeaf(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(8)
	b.InsertLeft(5, 8)
	b.InsertLeft(9, 5)
	b.InsertRight(7, 5)
	b.InsertLeft(1, 7)
	b.InsertRight(12, 7)
	b.InsertLeft(2, 12)
	b.InsertRight(4, 8)
	b.InsertRight(11, 4)
	b.InsertLeft(3, 11)

	assert.True(t, b.Find(9).IsLeaf())
	assert.True(t, b.Find(1).IsLeaf())
	assert.True(t, b.Find(3).IsLeaf())
	assert.False(t, b.Find(5).IsLeaf())
	assert.False(t, b.Find(12).IsLeaf())
	assert.False(t, b.Find(8).IsLeaf())
}

func TestPrintInverseLevelOrder(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(1)
	b.InsertLeft(2, 1)
	b.InsertRight(3, 1)
	b.InsertLeft(4, 2)
	b.InsertRight(5, 2)
	b.InsertLeft(6, 3)
	b.InsertRight(7, 3)

	assert.Equal(t, "[1 3 2 7 6 5 4]", b.PrintInverseLevelOrder())
}

func TestFindParent(t *testing.T) {
	b := binarytree.NewBinaryTree()
	b.SetRoot(1)
	b.InsertLeft(2, 1)
	b.InsertRight(3, 1)
	b.InsertLeft(4, 2)
	b.InsertRight(5, 2)
	b.InsertLeft(6, 3)
	b.InsertRight(7, 3)

	assert.Equal(t, 2, b.FindParent(4).Value)
	assert.Equal(t, 2, b.FindParent(5).Value)
	assert.Equal(t, 3, b.FindParent(6).Value)
	assert.Equal(t, 3, b.FindParent(7).Value)
	assert.Equal(t, 1, b.FindParent(2).Value)
	assert.Equal(t, 1, b.FindParent(3).Value)
	assert.Nil(t, b.FindParent(1))
}
