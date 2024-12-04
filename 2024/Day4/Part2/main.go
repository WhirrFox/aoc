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

	count := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if Get(x, y) != 'A' {
				continue
			}

			if IsSRight(x, y) {
				count++
			}
			if IsSLeft(x, y) {
				count++
			}
			if IsSBottom(x, y) {
				count++
			}
			if IsSTop(x, y) {
				count++
			}
		}
	}

	fmt.Println("Count:", count)
}

func IsSRight(x, y int) bool {
	return Get(x-1, y-1) == 'M' &&
		Get(x-1, y+1) == 'M' &&
		Get(x+1, y-1) == 'S' &&
		Get(x+1, y+1) == 'S'
}

func IsSLeft(x, y int) bool {
	return Get(x+1, y-1) == 'M' &&
		Get(x+1, y+1) == 'M' &&
		Get(x-1, y-1) == 'S' &&
		Get(x-1, y+1) == 'S'
}

func IsSBottom(x, y int) bool {
	return Get(x-1, y-1) == 'M' &&
		Get(x+1, y-1) == 'M' &&
		Get(x-1, y+1) == 'S' &&
		Get(x+1, y+1) == 'S'
}

func IsSTop(x, y int) bool {
	return Get(x-1, y+1) == 'M' &&
		Get(x+1, y+1) == 'M' &&
		Get(x-1, y-1) == 'S' &&
		Get(x+1, y-1) == 'S'
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
