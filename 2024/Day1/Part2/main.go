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

	slices.Sort(listB)

	score := 0
	for pointerA := 0; pointerA < len(listA); pointerA++ {
		target := listA[pointerA]
		pointerB, found := slices.BinarySearch(listB, target)
		if !found {
			continue
		}
		multiplier := 0
		for listB[pointerB] == target {
			multiplier++
			pointerB++
		}
		score += target * multiplier
	}
	fmt.Println("Similarity score: ", score)
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
