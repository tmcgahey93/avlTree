package main

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

func main() {
	// Create a new node
	node := &Node[int]{Value: 42}

	// Create another node and link it to the first one
	node2 := &Node[int]{Value: 84, Prev: node}
	node.Next = node2

	// Print the values of the nodes
	fmt.Println("First Node Value:", node.Value)
	fmt.Println("Second Node Value:", node.Next.Value)
	fmt.Println("Second Node's Previous Value:", node.Next.Prev.Value)
}
