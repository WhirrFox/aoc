package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var grid Grid
var goDownCache map[XY]int

func init() {
	goDownCache = make(map[XY]int)
}

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

	timelines := goDown(firstX, 1)

	fmt.Println("Timelines:", timelines)
}

func goDown(x, y int) int {
	switch grid.At(x, y) {
	case 'X':
		return 1
	case '.':
		return goDown(x, y+1)
	case '^':
		if _, ok := goDownCache[XY{x, y}]; ok {
			return goDownCache[XY{x, y}]
		}
		count := goDown(x-1, y) + goDown(x+1, y)
		goDownCache[XY{x, y}] = count
		return count
	default:
		panic("Character not supported")
	}
}
