package main

// this code doesn't work

import (
	"bufio"
	"fmt"
	"os"
)

type Garden []string
type Pos struct {
	X int
	Y int
}

var width, height = 0, 0

func main() {
	file, err := os.Open("example.txt")
	// file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var garden = Garden{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		garden = append(garden, scanner.Text())
	}
	width, height = len(garden[0]), len(garden)

	totalPrice := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if b := garden[y][x]; b != 0 {
				area := Area{
					Type: b,
				}
				totalPrice += area.CalculateArea(&garden, x, y)
			}
		}
	}

	fmt.Println("Sum of scores:", totalPrice)
}

func Get(m Garden, x, y int) byte {
	if x < 0 ||
		y < 0 ||
		x >= width ||
		y >= height {
		return 0
	}

	return m[y][x]
}

func Set(m Garden, x, y int, b byte) bool {
	if x < 0 ||
		y < 0 ||
		x >= width ||
		y >= height {
		return false
	}

	m[y] = replaceAtIndex(m[y], b, x)
	return true
}

func replaceAtIndex(in string, r byte, i int) string {
	out := []byte(in)
	out[i] = r
	return string(out)
}
