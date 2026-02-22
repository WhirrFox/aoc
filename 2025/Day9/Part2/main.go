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

type Line struct {
	Pos1, Pos2 XY
}

func (l *Line) CrossesLine(pos1, pos2 XY) bool {
	if l.Pos1.X == l.Pos2.X {
		// pos1.X < line < pos2.X
		return min(pos1.X, pos2.X) < l.Pos1.X && l.Pos1.X < max(pos1.X, pos2.X)
	} else if l.Pos1.Y == l.Pos2.Y {
		// pos1.Y < line < pos2.Y
		return min(pos1.Y, pos2.Y) < l.Pos1.Y && l.Pos1.Y < max(pos1.Y, pos2.Y)
	}
	panic("Not a straight line")
}

type XYMatch struct {
	Pos1, Pos2 XY
	Size       int
}

var positions = []XY{}
var lines = []Line{}

func main() {
	file, err := os.Open("example.txt")
	// file, err := os.Open("input.txt")
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
	for i := range len(positions) - 1 {
		lines = append(lines, Line{positions[i], positions[i+1]})
	}

	fmt.Println(SquareInsidePoly(XY{7, 3}, XY{11, 1}))
	return

	maxSize := XYMatch{XY{}, XY{}, math.MinInt}
	for i := range positions {
		for j := i + 1; j < len(positions); j++ {
			size := abs(positions[i].X-positions[j].X+1) * abs(positions[i].Y-positions[j].Y+1)
			if size > maxSize.Size && SquareInsidePoly(positions[i], positions[j]) {
				maxSize = XYMatch{positions[i], positions[j], size}
			}
		}
	}

	fmt.Println(maxSize.Pos1, maxSize.Pos2)
	fmt.Println("Largest size:", maxSize.Size)
}

func SquareInsidePoly(a, b XY) bool {
	p1 := XY{min(a.X, b.X), min(a.Y, b.Y)}
	p2 := XY{min(a.X, b.X), max(a.Y, b.Y)}
	p3 := XY{max(a.X, b.X), min(a.Y, b.Y)}
	p4 := XY{max(a.X, b.X), max(a.Y, b.Y)}
	for x := p1.X; x <= p3.X; x++ {
		if !InsidePoly(x, p1.Y) {
			return false
		}
	}
	for y := p1.Y; y <= p2.Y; y++ {
		if !InsidePoly(p1.X, y) {
			return false
		}
	}
	for x := p2.X; x <= p4.X; x++ {
		if !InsidePoly(x, p4.Y) {
			return false
		}
	}
	for y := p3.Y; y <= p4.Y; y++ {
		if !InsidePoly(p4.X, y) {
			return false
		}
	}
	return true
}

// https://www.geeksforgeeks.org/dsa/how-to-check-if-a-given-point-lies-inside-a-polygon/
func InsidePoly(x, y int) bool {
	num_vertices := len(positions)
	inside := false

	// Store the first point in the polygon and initialize the second point
	p1 := positions[0]

	// Loop through each edge in the polygon
	for i := 1; i < num_vertices+1; i++ {
		// Get the next point in the polygon
		p2 := positions[i%num_vertices]

		if p1.Y == p2.Y {
			if p1.X < p2.X {
				p1.X++
			} else {
				p1.X--
			}
		} else if p1.X == p2.X {
			if p1.Y < p2.Y {
				p1.Y++
			} else {
				p1.Y--
			}
		}

		// Check if the point is above the minimum y coordinate of the edge
		// Check if the point is below the maximum y coordinate of the edge
		// Check if the point is to the left of the maximum x coordinate of the edge
		if y >= min(p1.Y, p2.Y) && y <= max(p1.Y, p2.Y) && x <= max(p1.X, p2.X) {
			// Calculate the x-intersection of the line connecting the point to the edge
			x_intersection := math.MaxInt
			if p2.Y-p1.Y != 0 {
				x_intersection = (y-p1.Y)*(p2.X-p1.X)/(p2.Y-p1.Y) + p1.X
			}

			// Check if the point is on the same line as the edge or to the left of the x-intersection
			if p1.X == p2.X || x <= x_intersection {
				// Flip the inside flag
				inside = !inside
			}
		}
		// Store the current point as the first point for the next iteration
		p1 = p2
	}

	// Return the value of the inside flag
	return inside
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
