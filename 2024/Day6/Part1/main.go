package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LayoutMap []string

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
			guard = NewGuard(i, len(layoutMap))
			fmt.Println("Guard found!", guard)
		}
		layoutMap = append(layoutMap, scanner.Text())
	}
	width = len(layoutMap[0])
	height = len(layoutMap)

	for {
		if guard.X < 0 || guard.Y < 0 ||
			guard.X >= width || guard.Y >= height {
			break
		}

		layoutMap[guard.Y] = replaceAtIndex(layoutMap[guard.Y], 'X', guard.X)
		if guard.LookingAt() == '#' {
			guard.Rotate()
		} else {
			guard.Move()
		}
	}

	count := 0
	for _, s := range layoutMap {
		fmt.Println(s)
		count += strings.Count(s, "X")
	}
	fmt.Println("Distinct positions:", count)
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

func replaceAtIndex(in string, r byte, i int) string {
	out := []byte(in)
	out[i] = r
	return string(out)
}
