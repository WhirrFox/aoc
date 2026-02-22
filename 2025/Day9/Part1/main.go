package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type XY struct {
	X, Y int
}

type XYMatch struct {
	Pos1, Pos2 XY
	Size       int
}

var positions = []XY{}

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}
		positions = append(positions, XY{x, y})
	}

	max := XYMatch{XY{}, XY{}, math.MinInt}
	for i := range positions {
		for j := i + 1; j < len(positions); j++ {
			size := abs(positions[i].X-positions[j].X+1) * abs(positions[i].Y-positions[j].Y+1)
			if size > max.Size {
				max = XYMatch{positions[i], positions[j], size}
			}
		}
	}

	fmt.Println(max.Pos1, max.Pos2)
	fmt.Println("Largest size:", max.Size)
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
