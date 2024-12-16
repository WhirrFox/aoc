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

	m := MapMap{
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
			m.M = append(m.M, scanner.Text())
			if x := strings.Index(scanner.Text(), "@"); x != -1 {
				r.Pos.X = x
				r.Pos.Y = y
				m.Set(x, y, '.')
			}
			y++
		case 1:
			r.Moves += scanner.Text()
		}
	}
	m.Width, m.Height = len(m.M[0]), len(m.M)

	for {
		d := r.NextMove()
		if d == NONE {
			break
		}
		r.Move(&m, d)
	}

	for _, s := range m.M {
		fmt.Println(s)
	}
	fmt.Println(r.Pos)

	sum := 0
	for y, line := range m.M {
		for x, b := range line {
			if b == 'O' {
				sum += 100*y + x
			}
		}
	}
	fmt.Println("Sum:", sum)
}
