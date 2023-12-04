package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		sum += processCard(scanner.Text())
	}

	fmt.Println(sum)
}

func processCard(s string) int {
	sum := 0
	s = strings.Split(s, ":")[1]
	split := strings.Split(s, "|")
	winNum := strings.Split(split[0], " ")
	num := strings.Split(split[1], " ")

	for _, n := range num {
		if n == "" {
			continue
		}
		if slices.Contains(winNum, n) {
			if sum == 0 {
				sum = 1
			} else {
				sum *= 2
			}
		}
	}

	return sum
}
