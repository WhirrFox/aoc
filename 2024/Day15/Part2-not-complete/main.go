package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	origMap := MapMap{
		Width:  10000000,
		Height: 10000000,
	}
	r := Robot{}
	step := 0
	y := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			step++
			continue
		}

		switch step {
		case 0:
			line := ""
			for _, b := range scanner.Text() {
				switch b {
				case '#':
					line += "##"
				case 'O':
					line += "[]"
				case '.':
					line += ".."
				case '@':
					line += "@."
				}
			}
			origMap.M = append(origMap.M, line)
			if x := strings.Index(line, "@"); x != -1 {
				r.Pos.X = x
				r.Pos.Y = y
				origMap.Set(x, y, '.')
			}
			y++
		case 1:
			r.Moves += scanner.Text()
		}
	}
	origMap.Width, origMap.Height = len(origMap.M[0]), len(origMap.M)

	for {
		d := r.NextMove()
		if d == NONE {
			break
		}
		r.Move(&origMap, d)
	}

	for _, s := range origMap.M {
		fmt.Println(s)
	}
	fmt.Println(r.Pos)

	sum := 0
	for y, line := range origMap.M {
		for x, b := range line {
			if b == 'O' {
				sum += 100*y + x
			}
		}
	}
	fmt.Println("Sum:", sum)
}
