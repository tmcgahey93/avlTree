package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type node[T any] struct {
	Value      T
	Height     int
	Parent     *node[T]
	RightChild *node[T]
	LeftChild  *node[T]
}

func main() {

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

	var rootNode *node[int]

	for i := 0; i < len(numberSlice); i++ {

		rootNode = insertNode(rootNode, numberSlice[i])
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	searchNum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number")
	} else {
		fmt.Println("You entered:", searchNum)
	}

	//breadthSearch(searchSlice, searchNum)

	depthSearch(rootNode, searchNum)
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

// Recursive AVL insert
func insertNode[T int](n *node[T], value T) *node[T] {

	if n == nil {
		return &node[T]{Value: value, Height: 1}
	}

	if value < n.Value {
		n.LeftChild = insertNode(n.LeftChild, value)
	} else if value > n.Value {
		n.RightChild = insertNode(n.RightChild, value)
	} else {
		// Duplicate values not allowed
		return n
	}

	// Update height
	n.Height = 1 + max(height(n.LeftChild), height(n.RightChild))

	// Check balance
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

func breadthSearch[T constraints.Ordered](searchSlice []*node[T], searchNum T) {

	if len(searchSlice) == 0 {
		fmt.Printf("Breadth Search Value Not Found: %v\n", searchNum)
		return
	}

	if searchSlice[0].Value == searchNum {
		fmt.Println("Breadth Search Value Found")
	} else {

		fmt.Printf("Breadth Search Current Value: %v\n", searchSlice[0].Value)

		if searchSlice[0].LeftChild != nil {
			searchSlice = append(searchSlice, searchSlice[0].LeftChild)
		}

		if searchSlice[0].RightChild != nil {
			searchSlice = append(searchSlice, searchSlice[0].RightChild)
		}

		searchSlice = searchSlice[1:]
		breadthSearch(searchSlice, searchNum)
	}
}

func depthSearch[T constraints.Ordered](node *node[T], value T) {

	fmt.Printf("Depth Search Current Value: %v\n", node.Value)

	if node.Value == value {
		fmt.Printf("Depth Search Value Found!: %v\n", value)
		return //need to return search path
	} else if value < node.Value {
		//Traverse left branch
		if node.LeftChild != nil {
			depthSearch(node.LeftChild, value)
		} else {
			//Not Found
			fmt.Println("Depth Search Value Not Found")
			return
		}
	} else {
		//Travere right branch
		if node.RightChild != nil {
			depthSearch(node.RightChild, value)
		} else {
			//Not Found
			fmt.Println("Depth Search Value Not Found")
			return
		}
	}
}

func updateHeight[T constraints.Ordered](n *node[T]) {
	if n != nil {
		n.Height = 1 + max(height(n.LeftChild), height(n.RightChild))
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func height[T constraints.Ordered](n *node[T]) int {
	if n == nil {
		return 0
	}
	return n.Height
}

// Get balance factor
func getBalance[T int](n *node[T]) int {
	if n == nil {
		return 0
	}
	return height(n.LeftChild) - height(n.RightChild)
}

//		    z
//		   /
//		  y
//		 /
//	    x
//
// x := z.LeftChild.LeftChild
// y := z.LeftChild
// z := unbalancedNode
// Right rotation
// Perform a right rotation
func rightRotation[T int](z *node[T]) *node[T] {
	y := z.LeftChild
	T2 := y.RightChild

	// Perform rotation
	y.RightChild = z
	z.LeftChild = T2

	// Update heights
	updateHeight(z)
	updateHeight(y)

	// Return new root
	return y
}

// RR leftRotaion
//
//		   z
//		    \
//		     y
//		      \
//	           x
//
// x := z.RightChild.RightChild
// y := z.RightChild
// z := unbalancedNode
func leftRotation[T int](z *node[T]) *node[T] {
	y := z.RightChild
	T2 := y.LeftChild

	// Perform rotation
	y.LeftChild = z
	z.RightChild = T2

	// Update heights
	updateHeight(z)
	updateHeight(y)

	// Return new root
	return y
}
