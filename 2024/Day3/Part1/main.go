package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	content := ""
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		content += scanner.Text()
	}

	sum := 0
	re := regexp.MustCompile(`(?m)mul\(\d+,\d+\)`)
	for _, match := range re.FindAllString(content, -1) {
		nums := strings.Split(match[4:len(match)-1], ",")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		sum += num1 * num2
	}
	fmt.Println("Sum:", sum)
}
