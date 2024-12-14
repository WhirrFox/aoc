package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, err := os.Open("example.txt")
	// m := NewMap(11, 7)

	file, err := os.Open("input.txt")
	m := NewMap(101, 103)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var robots = []Robot{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		robots = append(robots, NewRobot(scanner.Text()))
	}

	quads := make(map[int]int)
	for _, r := range robots {
		r.Move(&m, 100)
		fmt.Println(r)
		quads[r.GetQuadrant(&m)]++
	}

	for key, value := range quads {
		fmt.Println(key, value)
	}
	fmt.Println("You have to manually multiply these numbers (except -1).")
	fmt.Println("I was just too lazy to implement that, since I had to debug a error for a long time...")
}
