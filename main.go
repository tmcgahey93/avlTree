package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Node[T any] struct {
	Value         T
	BalanceFactor int
	Parent        *Node[T]
	RightChild    *Node[T]
	LeftChild     *Node[T]
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
}

func readFile(filename string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
