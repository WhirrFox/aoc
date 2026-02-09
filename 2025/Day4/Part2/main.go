package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid []string
type XY struct {
	x, y int
}

func (g *Grid) At(x, y int) string {
	if x < 0 || x >= len((*g)[0]) || y < 0 || y >= len((*g)) {
		return "."
	}
	return string((*g)[y][x])
}

func (g *Grid) AtPos(pos XY) string {
	return g.At(pos.x, pos.y)
}

func (g *Grid) Set(x, y int, s string) {
	runeSlice := []rune((*g)[y])
	runeSlice[x] = rune(s[0])
	(*g)[y] = string(runeSlice)
}

func (g *Grid) SetAtPos(pos XY, s string) {
	g.Set(pos.x, pos.y, s)
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

	removablePaper := 0
	for {
		canBePickedLocations := []XY{}
		for y := range len(grid) {
			for x := range len(grid[0]) {
				if canPickedUp(x, y) {
					canBePickedLocations = append(canBePickedLocations, XY{x, y})
				}
			}
		}

		if len(canBePickedLocations) == 0 {
			break
		}
		removablePaper += len(canBePickedLocations)

		for _, p := range canBePickedLocations {
			grid.SetAtPos(p, ".")
		}
	}

	fmt.Println("Removable paper:", removablePaper)
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
