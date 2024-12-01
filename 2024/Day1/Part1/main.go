package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	listA := []int{}
	listB := []int{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		left, right := getNumbers(scanner.Text())
		listA = append(listA, left)
		listB = append(listB, right)
	}

	slices.Sort(listA)
	slices.Sort(listB)

	distance := 0
	for i := 0; i < len(listA); i++ {
		distance += abs(listA[i] - listB[i])
	}
	fmt.Println("Distance: ", distance)
}

func getNumbers(text string) (int, int) {
	numbers := strings.Split(text, "   ")
	left, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(err)
	}
	return left, right
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
