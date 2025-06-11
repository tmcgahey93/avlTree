package avl

import (
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	var root *Node[int]
	values := []int{10, 20, 5, 4, 15, 25, 2}

	// Insert values
	for _, v := range values {
		root = Insert(root, v)
	}

	// Search values we inserted
	for _, v := range values {
		if node, found := DepthSearch(root, v); !found || node == nil || node.Value != v {
			t.Errorf("Expected to find %d, but did not", v)
		}
	}

	// Search for a value that doesn't exist
	if _, found := DepthSearch(root, 999); found {
		t.Errorf("Expected not to find 999, but did")
	}
}

func TestAVLBalance(t *testing.T) {
	var root *Node[int]

	// Insert values that should trigger rotations
	values := []int{30, 20, 10} // triggers right rotation

	root = nil
	for _, v := range values {
		root = Insert(root, v)
	}

	if root.Value != 20 {
		t.Errorf("Expected root value 20 after right rotation, got %d", root.Value)
	}

	if root.LeftChild == nil || root.LeftChild.Value != 10 {
		t.Errorf("Expected left child 10, got %+v", root.LeftChild)
	}

	if root.RightChild == nil || root.RightChild.Value != 30 {
		t.Errorf("Expected right child 30, got %+v", root.RightChild)
	}
}
