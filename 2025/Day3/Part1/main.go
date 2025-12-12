package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		highest := 0
		for a := range len(line) {
			for b := a + 1; b < len(line); b++ {
				num, err := strconv.Atoi(string(line[a]) + string(line[b]))
				if err != nil {
					panic(err)
				}
				if num > highest {
					highest = num
				}
			}
		}
		sum += highest
	}
	fmt.Println("Sum:", sum)
}
