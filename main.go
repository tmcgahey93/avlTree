package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Node[T any] struct {
	Value      T
	Height     int
	Parent     *Node[T]
	RightChild *Node[T]
	LeftChild  *Node[T]
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

	var rootNode *Node[int]

	for i := 0; i < len(numberSlice); i++ {
		node := &Node[int]{Value: numberSlice[i], Height: 0}
		fmt.Println("Node created with value:", node.Value)

		if i == 0 {
			rootNode = node
			continue
		}

		rootNode = insertNode(rootNode, node)
		fmt.Println("rootNode: ", rootNode.Value)
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

func insertNode[T int](currentNode *Node[T], newNode *Node[T]) *Node[T] {
	if newNode.Value < currentNode.Value {
		if currentNode.LeftChild == nil {
			currentNode.LeftChild = newNode
			newNode.Parent = currentNode
		} else {
			insertNode(currentNode.LeftChild, newNode)
		}
	} else if currentNode.RightChild == nil {
		currentNode.RightChild = newNode
		newNode.Parent = currentNode
	} else {
		insertNode(currentNode.RightChild, newNode)
	}

	//Need to set height when returning from recursive call
	updateHeight(currentNode)

	balanceFactor := balanceFactor(currentNode)

	//Out of balance on left
	if balanceFactor < -1 {
		//LL - Single Right Rotation
		if newNode.Value < currentNode.LeftChild.Value {
			fmt.Println("Rotating Right")
			currentNode = rightRotation(currentNode)
			fmt.Println("Root Node returning from function: ", currentNode.Value)
		} else {

		}

		//Out of balance on right
	} else if balanceFactor > 1 {

		//RR - Single Left Rotation
		if newNode.Value > currentNode.RightChild.Value {
			fmt.Println("Rotating Left")
			leftRotation(currentNode)
		} else {

		}

	}

	return currentNode
}

func breadthSearch[T constraints.Ordered](searchSlice []*Node[T], searchNum T) {

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

func depthSearch[T constraints.Ordered](currentNode *Node[T], searchNum T) {

	fmt.Printf("Depth Search Current Value: %v\n", currentNode.Value)

	if currentNode.Value == searchNum {
		fmt.Printf("Depth Search Value Found!: %v\n", searchNum)
		return //need to return search path
	} else if searchNum < currentNode.Value {
		//Traverse left branch
		if currentNode.LeftChild != nil {
			depthSearch(currentNode.LeftChild, searchNum)
		} else {
			//Not Found
			fmt.Println("Depth Search Value Not Found")
			return
		}
	} else {
		//Travere right branch
		if currentNode.RightChild != nil {
			depthSearch(currentNode.RightChild, searchNum)
		} else {
			//Not Found
			fmt.Println("Depth Search Value Not Found")
			return
		}
	}
}

func updateHeight[T constraints.Ordered](n *Node[T]) {
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

func height[T constraints.Ordered](n *Node[T]) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func balanceFactor[T int](n *Node[T]) int {
	if n == nil {
		return 0
	}
	fmt.Println("Calc BF: ", n.Value)

	if n.RightChild == nil {
		return (-1 - height(n.LeftChild))
	} else if n.LeftChild == nil {
		return (1 + height(n.RightChild))
	} else {
		return height(n.RightChild) - height(n.LeftChild)
	}
}

func rightRotation[T int](currentNode *Node[T]) *Node[T] {

	holdNode := currentNode

	//Not the root node
	if currentNode.Parent != nil {
		if currentNode.Parent.RightChild == currentNode {
			currentNode.Parent.RightChild = currentNode.LeftChild
		} else {
			currentNode.Parent.LeftChild = currentNode.LeftChild
		}
	} else {
		//Need to make left child the new root node
		//Current Node becomes the right child of the new root node
		//The new root node still has its same left child
		currentNode = holdNode.LeftChild
		currentNode.RightChild = holdNode
		holdNode.Parent = currentNode
	}

	return currentNode
}

func leftRotation[T int](currentNode *Node[T]) *Node[T] {

	holdNode := currentNode

	//Not the root node
	if currentNode.Parent != nil {
		if currentNode.Parent.RightChild == currentNode {
			currentNode.Parent.RightChild = currentNode.LeftChild
		} else {
			currentNode.Parent.LeftChild = currentNode.LeftChild
		}
	} else {
		//Need to make left child the new root node
		//Current Node becomes the right child of the new root node
		//The new root node still has its same left child
		currentNode = holdNode.LeftChild
		currentNode.RightChild = holdNode
		holdNode.Parent = currentNode
	}

	return currentNode

}
