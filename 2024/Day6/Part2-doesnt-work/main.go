package main

// this code doesn't work

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type LayoutMap []string

var origMap = LayoutMap{}
var layoutMap = LayoutMap{}
var width, height = 0, 0
var guard = NewGuard(0, 0)

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if i := strings.Index(scanner.Text(), "^"); i != -1 {
			guard = NewGuard(i, len(origMap))
		}
		origMap = append(origMap, scanner.Text())
	}
	width = len(origMap[0])
	height = len(origMap)

	layoutMap = slices.Clone(origMap)
	guard.Move()

	steps := 0
	for {
		steps++
		fmt.Println("step", steps)
		if guard.X < 0 || guard.Y < 0 ||
			guard.X >= width || guard.Y >= height {
			break
		}

		guard.SetInFront('#')
		newGuard := guard
		if newGuard.IsStuck() {
			fmt.Println("\tStuck!")
			layoutMap = slices.Clone(origMap)
			guard.SetInFront('O')
			origMap = slices.Clone(layoutMap)
		} else {
			layoutMap = slices.Clone(origMap)
		}

		if guard.LookingAt() == '#' {
			guard.Rotate()
		} else {
			guard.Move()
		}
	}

	count := 0
	for _, s := range origMap {
		fmt.Println(s)
		count += strings.Count(s, "O")
	}
	fmt.Println("Possible obstacles:", count)
}

func Get(x, y int) byte {
	if x < 0 ||
		y < 0 ||
		x >= width ||
		y >= height {
		return 0
	}

	return layoutMap[y][x]
}

func Set(x, y int, b byte) {
	if x < 0 ||
		y < 0 ||
		x >= width ||
		y >= height {
		return
	}

	layoutMap[y] = replaceAtIndex(layoutMap[y], b, x)
}

func replaceAtIndex(in string, r byte, i int) string {
	out := []byte(in)
	out[i] = r
	return string(out)
}

func PrintLayout() {
	for _, s := range layoutMap {
		fmt.Println(s)
	}
}
