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

	searchSlice := []*Node[int]{rootNode}

	breadthSearch(searchSlice, searchNum)

	//depthSearch(rootNode, searchNum)
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

func breadthSearch[T constraints.Ordered](searchSlice []*Node[T], searchNum T) {

	if len(searchSlice) == 0 {
		fmt.Println("Breadth Search Value Not Found")
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

	fmt.Println(currentNode.Value)

	if currentNode.Value == searchNum {
		fmt.Println("Depth Search Value Found")
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
