package main

import (
	"bufio"
	"fmt"
	"os"
)

type HikingMap []string
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

	var hikingMap = HikingMap{}
	var trailheads = []Trailhead{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		for i, b := range scanner.Text() {
			if b == '0' {
				trailheads = append(trailheads, NewTrailhead(i, len(hikingMap)))
			}
		}
		hikingMap = append(hikingMap, scanner.Text())
	}
	width, height = len(hikingMap[0]), len(hikingMap)

	scoreSum := 0
	for _, t := range trailheads {
		score := t.CalculateScore(&hikingMap)
		fmt.Println(score)
		scoreSum += score
	}

	fmt.Println("Sum of scores:", scoreSum)
}

func Get(m HikingMap, x, y int) byte {
	if x < 0 ||
		y < 0 ||
		x >= width ||
		y >= height {
		return 0
	}

	return m[y][x]
}
