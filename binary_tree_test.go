package data_structure_test

import (
	"testing"

	data_structure "github.com/ianadiwibowo/central-park"
)

func TestBinaryTreeNewBinaryTree(t *testing.T) {
	b := data_structure.NewBinaryTree()

	if b.PrintPreOrder() != "[]" {
		t.Errorf("Expected: []. Got: %v", b.PrintPreOrder())
	}
}

func TestBinaryTreeSetRoot(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(5)

	if b.PrintPreOrder() != "[5]" {
		t.Errorf("Expected: [5]. Got: %v", b.PrintPreOrder())
	}
}

func TestBinaryTreeInsertLeft(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(5)
	b.InsertLeft(1, 5)

	if b.PrintPreOrder() != "[5 1]" {
		t.Errorf("Expected: [5 1]. Got: %v", b.PrintPreOrder())
	}
}

func TestBinaryTreeInsertLeftOnNonExistentValue(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(5)
	b.InsertLeft(1, 4)

	if b.PrintPreOrder() != "[5]" {
		t.Errorf("Expected: [5]. Got: %v", b.PrintPreOrder())
	}
}

func TestBinaryTreeInsertRight(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(5)
	b.InsertRight(7, 5)

	if b.PrintPreOrder() != "[5 7]" {
		t.Errorf("Expected: [5 7]. Got: %v", b.PrintPreOrder())
	}
}

func TestBinaryTreeInsertRightOnNonExistentValue(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(5)
	b.InsertRight(7, 4)

	if b.PrintPreOrder() != "[5]" {
		t.Errorf("Expected: [5]. Got: %v", b.PrintPreOrder())
	}
}

func TestBinaryTreeFind(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(5)
	b.InsertRight(7, 5)
	currentNode := b.Find(7)

	if currentNode.Value != 7 {
		t.Errorf("Expected: 7. Got: %v", currentNode.Value)
	}
}

func TestBinaryTreeFindNonExistentValue(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(5)
	b.InsertRight(7, 5)
	currentNode := b.Find(8)

	if currentNode != nil {
		t.Errorf("Expected: nil. Got: %v", currentNode)
	}
}

func TestBinaryTreeHeightOnBalancedTree(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(5)
	b.InsertLeft(1, 5)
	b.InsertRight(2, 5)

	if b.Height() != 1 {
		t.Errorf("Expected: 1. Got: %v", b.Height())
	}
}

func TestBinaryTreeHeightOnBiggerBalancedTree(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(1)
	b.InsertLeft(2, 1)
	b.InsertLeft(3, 2)
	b.InsertRight(4, 2)
	b.InsertRight(5, 1)
	b.InsertLeft(6, 5)
	b.InsertRight(7, 5)

	if b.Height() != 2 {
		t.Errorf("Expected: 2. Got: %v", b.Height())
	}
}

func TestBinaryTreeHeightOnSkewedLeftTree(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(1)
	b.InsertLeft(2, 1)
	b.InsertLeft(3, 2)
	b.InsertLeft(4, 3)
	b.InsertLeft(5, 4)

	if b.Height() != 4 {
		t.Errorf("Expected: 4. Got: %v", b.Height())
	}
}

func TestBinaryTreeHeightOnSkewedRightTree(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(1)
	b.InsertRight(2, 1)
	b.InsertRight(3, 2)
	b.InsertRight(4, 3)
	b.InsertRight(5, 4)

	if b.Height() != 4 {
		t.Errorf("Expected: 4. Got: %v", b.Height())
	}
}

func TestBinaryTreeHeightOnRootOnlyTree(t *testing.T) {
	b := data_structure.NewBinaryTree()
	b.SetRoot(1)

	if b.Height() != 0 {
		t.Errorf("Expected: 0. Got: %v", b.Height())
	}
}

func TestBinaryTreeHeightOnEmptyTree(t *testing.T) {
	b := data_structure.NewBinaryTree()

	if b.Height() != 0 {
		t.Errorf("Expected: 0. Got: %v", b.Height())
	}
}

func TestBinaryPrintPreOrder(t *testing.T) {
	b := data_structure.NewBinaryTree()
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

	if b.PrintPreOrder() != "[8 5 9 7 1 12 2 4 11 3]" {
		t.Errorf("Expected: [8 5 9 7 1 12 2 4 11 3]. Got: %v", b.PrintPreOrder())
	}
}

func TestBinaryTreePrintInOrder(t *testing.T) {
	b := data_structure.NewBinaryTree()
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

	if b.PrintInOrder() != "[9 5 1 7 2 12 8 4 3 11]" {
		t.Errorf("Expected: [9 5 1 7 2 12 8 4 3 11]. Got: %v", b.PrintInOrder())
	}
}

func TestBinaryTreePrintPostOrder(t *testing.T) {
	b := data_structure.NewBinaryTree()
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

	if b.PrintPostOrder() != "[9 1 2 12 7 5 3 11 4 8]" {
		t.Errorf("Expected: [9 1 2 12 7 5 3 11 4 8]. Got: %v", b.PrintPostOrder())
	}
}

// func TestBinaryTreePrintLevelOrder(t *testing.T) {
// 	b := data_structure.NewBinaryTree()
// 	b.SetRoot(8)
// 	b.InsertLeft(5, 8)
// 	b.InsertLeft(9, 5)
// 	b.InsertRight(7, 5)
// 	b.InsertLeft(1, 7)
// 	b.InsertRight(12, 7)
// 	b.InsertLeft(2, 12)
// 	b.InsertRight(4, 8)
// 	b.InsertRight(11, 4)
// 	b.InsertLeft(3, 11)

// 	if b.PrintLevelOrder() != "[8 5 4 9 7 11 1 12 3 2]" {
// 		t.Errorf("Expected: [8 5 4 9 7 11 1 12 3 2]. Got: %v", b.PrintLevelOrder())
// 	}
// }
