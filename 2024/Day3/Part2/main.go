package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
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

	dos := getDos(&content)
	donts := getDonts(&content)

	sum := 0
	re := regexp.MustCompile(`(?m)mul\(\d+,\d+\)`)
	for _, matchRange := range re.FindAllStringIndex(content, -1) {
		match := content[matchRange[0]:matchRange[1]]

		DoIndex, _ := slices.BinarySearch(dos, matchRange[0])
		prevDo := dos[DoIndex-1]
		DontIndex, _ := slices.BinarySearch(donts, matchRange[0])
		prevDont := donts[DontIndex-1]

		if prevDont > prevDo {
			continue
		}

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

func getDos(content *string) []int {
	dos := []int{-10, 0}
	re := regexp.MustCompile(`(?m)do\(\)`)
	for _, s := range re.FindAllStringIndex(*content, -1) {
		dos = append(dos, s[0])
	}
	return dos
}

func getDonts(content *string) []int {
	donts := []int{-5, -1}
	re := regexp.MustCompile(`(?m)don't\(\)`)
	for _, s := range re.FindAllStringIndex(*content, -1) {
		donts = append(donts, s[0])
	}
	return donts
}
