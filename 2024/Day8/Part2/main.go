package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MapMap []string
type Pos struct {
	X int
	Y int
}

var width, height = 0, 0

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	antennaMap := MapMap{}
	antennas := make(map[rune][]Pos)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		antennaMap = append(antennaMap, scanner.Text())
	}
	width, height = len(antennaMap[0]), len(antennaMap)

	for y, line := range antennaMap {
		for x, c := range line {
			if c == '.' {
				continue
			}

			if _, ok := antennas[c]; !ok {
				antennas[c] = []Pos{}
			}
			antennas[c] = append(antennas[c], Pos{X: x, Y: y})
		}
	}

	for _, s := range antennas {
		for _, p1 := range s {
			for _, p2 := range s {
				if p1 == p2 {
					continue
				}

				diffX := p2.X - p1.X
				diffY := p2.Y - p1.Y

				antinode := Pos{
					X: p1.X + diffX*2,
					Y: p1.Y + diffY*2,
				}

				for Set(antennaMap, antinode.X, antinode.Y, '#') {
					antinode.X += diffX
					antinode.Y += diffY
				}
			}
		}

		if len(s) > 1 {
			for _, p := range s {
				Set(antennaMap, p.X, p.Y, '#')
			}
		}
	}

	PrintLayout(antennaMap)

	count := 0
	for _, l := range antennaMap {
		count += strings.Count(l, "#")
	}
	fmt.Println("Unique locations:", count)
}

func Get(m MapMap, x, y int) byte {
	if x < 0 ||
		y < 0 ||
		x >= width ||
		y >= height {
		return 0
	}

	return m[y][x]
}

func Set(m MapMap, x, y int, b byte) bool {
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

func PrintLayout(m MapMap) {
	for _, s := range m {
		fmt.Println(s)
	}
}
