package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	X int
	Y int
}

func main() {
	// file, err := os.Open("example.txt")
	// m := NewMap(7, 7)
	file, err := os.Open("input.txt")
	m := NewMap(71, 71)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	fallingBytes := []Pos{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		values := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		fallingBytes = append(fallingBytes, Pos{x, y})
	}
	m.Width, m.Height = len(m.M[0]), len(m.M)

	for i, p := range fallingBytes {
		fmt.Println(i)
		m.Set(p.X, p.Y, '#')
		m.CreateMap2()
		m.Spread(0, 0, EAST, 1)

		if m.Get2(m.Width-1, m.Height-1) == 0 {
			fmt.Println(p)
			break
		}
	}
}
