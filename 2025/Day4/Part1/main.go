package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid []string

func (g *Grid) At(x, y int) string {
	if x < 0 || x >= len((*g)[0]) || y < 0 || y >= len((*g)) {
		return "."
	}
	return string((*g)[y][x])
}

var grid Grid = []string{}

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	pickedPaper := 0
	for y := range len(grid) {
		for x := range len(grid[0]) {
			if canPickedUp(x, y) {
				pickedPaper++
			}
		}
	}

	fmt.Println("Picked up paper:", pickedPaper)
}

var locations = [][]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func canPickedUp(x, y int) bool {
	if grid.At(x, y) == "." {
		return false
	}

	paper := 0
	for _, l := range locations {
		if grid.At(x+l[0], y+l[1]) == "@" {
			paper++
		}
	}
	return paper < 4
}
