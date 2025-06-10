package avl

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	Value      T
	Height     int
	Parent     *Node[T]
	LeftChild  *Node[T]
	RightChild *Node[T]
}

func Insert[T constraints.Ordered](n *Node[T], value T) *Node[T] {
	if n == nil {
		return &Node[T]{Value: value, Height: 1}
	}

	if value < n.Value {
		n.LeftChild = Insert(n.LeftChild, value)
		n.LeftChild.Parent = n
	} else if value > n.Value {
		n.RightChild = Insert(n.RightChild, value)
		n.RightChild.Parent = n
	} else {
		return n // duplicate, do nothing
	}

	updateHeight(n)
	balance := getBalance(n)

	// Left Left
	if balance > 1 && value < n.LeftChild.Value {
		return rightRotation(n)
	}
	// Right Right
	if balance < -1 && value > n.RightChild.Value {
		return leftRotation(n)
	}
	// Left Right
	if balance > 1 && value > n.LeftChild.Value {
		n.LeftChild = leftRotation(n.LeftChild)
		return rightRotation(n)
	}
	// Right Left
	if balance < -1 && value < n.RightChild.Value {
		n.RightChild = rightRotation(n.RightChild)
		return leftRotation(n)
	}

	return n
}

func DepthSearch[T constraints.Ordered](n *Node[T], value T) (*Node[T], bool) {
	if n == nil {
		return nil, false
	}
	if value == n.Value {
		return n, true
	} else if value < n.Value {
		return DepthSearch(n.LeftChild, value)
	} else {
		return DepthSearch(n.RightChild, value)
	}
}

func updateHeight[T constraints.Ordered](n *Node[T]) {
	n.Height = 1 + max(height(n.LeftChild), height(n.RightChild))
}

func height[T constraints.Ordered](n *Node[T]) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBalance[T constraints.Ordered](n *Node[T]) int {
	if n == nil {
		return 0
	}
	return height(n.LeftChild) - height(n.RightChild)
}

func rightRotation[T constraints.Ordered](z *Node[T]) *Node[T] {
	y := z.LeftChild
	T2 := y.RightChild

	y.RightChild = z
	z.LeftChild = T2

	if T2 != nil {
		T2.Parent = z
	}
	y.Parent = z.Parent
	z.Parent = y

	updateHeight(z)
	updateHeight(y)

	return y
}

func leftRotation[T constraints.Ordered](z *Node[T]) *Node[T] {
	y := z.RightChild
	T2 := y.LeftChild

	y.LeftChild = z
	z.RightChild = T2

	if T2 != nil {
		T2.Parent = z
	}
	y.Parent = z.Parent
	z.Parent = y

	updateHeight(z)
	updateHeight(y)

	return y
}
