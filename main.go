package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"yourmodule/avltree/avl"
	"yourmodule/avltree/util"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file_name>")
		return
	}

	numbers, err := util.ReadNumbersFromFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var root *avl.Node[int]
	for _, num := range numbers {
		root = avl.Insert(root, num)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number to search: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	target, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if _, found := avl.DepthSearch(root, target); found {
		fmt.Println("Value found!")
	} else {
		fmt.Println("Value not found")
	}
}
