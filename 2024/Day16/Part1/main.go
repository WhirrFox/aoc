package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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

	fmt.Println("Sum:", m.Get2(m.Width-2, 1)-1)
}
