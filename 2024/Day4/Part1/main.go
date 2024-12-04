package main

import (
	"bufio"
	"fmt"
	"os"
)

type CharacterMap []string

var charMap = CharacterMap{}
var width, height = 0, 0

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

		charMap = append(charMap, scanner.Text())
	}
	width = len(charMap[0])
	height = len(charMap)

	count := RunThroughMap()

	fmt.Println("Count:", count)
}

func RunThroughMap() int {
	count := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if IsRight(x, y) {
				count++
			}
			if IsLeft(x, y) {
				count++
			}
			if IsBottomRight(x, y) {
				count++
			}
			if IsTopLeft(x, y) {
				count++
			}
			if IsBottom(x, y) {
				count++
			}
			if IsTop(x, y) {
				count++
			}
			if IsBottomLeft(x, y) {
				count++
			}
			if IsTopRight(x, y) {
				count++
			}
		}
	}
	return count
}

func IsRight(x, y int) bool {
	return Get(x, y) == 'X' &&
		Get(x+1, y) == 'M' &&
		Get(x+2, y) == 'A' &&
		Get(x+3, y) == 'S'
}

func IsLeft(x, y int) bool {
	return Get(x, y) == 'X' &&
		Get(x-1, y) == 'M' &&
		Get(x-2, y) == 'A' &&
		Get(x-3, y) == 'S'
}

func IsBottomRight(x, y int) bool {
	return Get(x, y) == 'X' &&
		Get(x+1, y+1) == 'M' &&
		Get(x+2, y+2) == 'A' &&
		Get(x+3, y+3) == 'S'
}

func IsTopLeft(x, y int) bool {
	return Get(x, y) == 'X' &&
		Get(x-1, y-1) == 'M' &&
		Get(x-2, y-2) == 'A' &&
		Get(x-3, y-3) == 'S'
}

func IsBottom(x, y int) bool {
	return Get(x, y) == 'X' &&
		Get(x, y+1) == 'M' &&
		Get(x, y+2) == 'A' &&
		Get(x, y+3) == 'S'
}

func IsTop(x, y int) bool {
	return Get(x, y) == 'X' &&
		Get(x, y-1) == 'M' &&
		Get(x, y-2) == 'A' &&
		Get(x, y-3) == 'S'
}

func IsBottomLeft(x, y int) bool {
	return Get(x, y) == 'X' &&
		Get(x-1, y+1) == 'M' &&
		Get(x-2, y+2) == 'A' &&
		Get(x-3, y+3) == 'S'
}

func IsTopRight(x, y int) bool {
	return Get(x, y) == 'X' &&
		Get(x+1, y-1) == 'M' &&
		Get(x+2, y-2) == 'A' &&
		Get(x+3, y-3) == 'S'
}

func Get(x, y int) byte {
	if x < 0 ||
		y < 0 ||
		x >= width ||
		y >= height {
		return 0
	}

	return charMap[y][x]
}
