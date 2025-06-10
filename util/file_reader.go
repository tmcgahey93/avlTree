package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadNumbersFromFile(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers, scanner.Err()
}
