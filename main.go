package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node[T any] struct {
	Value      T
	Parent     *Node[T]
	RightChild *Node[T]
	LeftChild  *Node[T]
}

func main() {
	// Create a new node

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run avlTree.go <file_name>")
		return
	}

	// Parse size from command-line argument
	fileName := os.Args[1]

	numberSlice, err := readFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(numberSlice)

	var rootNode *Node[int]

	for i := 0; i < len(numberSlice); i++ {
		node := &Node[int]{Value: numberSlice[i]}
		fmt.Println("Node created with value:", node.Value)

		if i == 0 {
			rootNode = node
			continue
		}

		insertNode(rootNode, node)

	}

	fmt.Println("Root Node Value:", rootNode.Value)
}

func readFile(filename string) ([]int, error) {

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var numbers []int

	// Create a scanner to read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Invalid number:", line)
			continue
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return numbers, nil
}

func insertNode[T int](root *Node[T], newNode *Node[T]) {
	if newNode.Value < root.Value {
		if root.LeftChild == nil {
			root.LeftChild = newNode
			newNode.Parent = root
		} else {
			root = root.LeftChild
			insertNode(root, newNode)
		}
	} else if root.RightChild == nil {
		root.RightChild = newNode
		newNode.Parent = root
	} else {
		root = root.RightChild
		insertNode(root, newNode)
	}
}
