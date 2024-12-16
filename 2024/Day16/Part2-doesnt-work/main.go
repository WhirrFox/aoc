package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

func main() {
	debug.SetMaxStack(4000000000)
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	m := MapMap{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		m.M = append(m.M, scanner.Text())
	}
	m.Width, m.Height = len(m.M[0]), len(m.M)

	m.CreateMap2()
	m.Spread(1, m.Height-2, EAST, 1)
	m.ReverseSpread(m.Width-2, 1)

	count := 0
	for _, s := range m.M {
		fmt.Println(s)
		count += strings.Count(s, "O")
	}
	fmt.Println("Tiles of best paths:", count)
}
