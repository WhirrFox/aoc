package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var rules = make(map[int][]int)
var pagesProduce = [][]int{}

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	step := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			step++
			continue
		}

		if step == 0 {
			nums := strings.Split(scanner.Text(), "|")
			x, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}

			_, ok := rules[x]
			if !ok {
				rules[x] = []int{y}
			} else {
				rules[x] = append(rules[x], y)
			}
		} else if step == 1 {
			s := strings.Split(scanner.Text(), ",")
			pages := make([]int, len(s))
			for i := range s {
				p, err := strconv.Atoi(s[i])
				if err != nil {
					panic(err)
				}
				pages[i] = p
			}
			pagesProduce = append(pagesProduce, pages)
		}
	}

	sum := 0
	for _, pages := range pagesProduce {
		if !IsPageOk(pages) {
			fixedPage := FixPages(slices.Clone(pages))
			sum += fixedPage[len(fixedPage)/2]
		}
	}

	fmt.Println("Sum:", sum)
}

func IsPageOk(pages []int) bool {
	for i, p := range pages {
		if i == 0 {
			continue
		}

		for _, r := range rules[p] {
			if slices.Contains(pages[:i], r) {
				return false
			}
		}
	}
	return true
}

func FixPages(pages []int) []int {
	for range len(pages) {
		for i, p := range pages {
			if i == 0 {
				continue
			}

			for _, r := range rules[p] {
				if i2 := slices.Index(pages[:i], r); i2 != -1 {
					pages[i], pages[i2] = pages[i2], pages[i]
				}
			}
		}
	}
	return pages
}
