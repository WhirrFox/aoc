package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var grid Grid
var splits = 0

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	firstX := -1
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
		if firstX < 0 {
			firstX = strings.Index(scanner.Text(), "S")
		}
	}

	goDown(firstX, 1)

	for _, r := range grid {
		fmt.Println(string(r))
	}
	fmt.Println("Splits:", splits)
}

func goDown(x, y int) {
	switch grid.At(x, y) {
	case 'X':
	case '|':
		return
	case '.':
		grid.Set(x, y, '|')
		goDown(x, y+1)
	case '^':
		splits++
		goDown(x-1, y)
		goDown(x+1, y)
	default:
		panic("Character not supported")
	}
}
