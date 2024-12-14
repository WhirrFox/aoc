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
	mapmap := NewMap(101, 103)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var robots = []*Robot{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		r := NewRobot(scanner.Text())
		robots = append(robots, &r)
	}

	for i := range 10000 {
		if i == 0 {
			continue
		}

		m := mapmap.Clone()
		for _, r := range robots {
			r.Move(&m, 1)
		}
		if m.Qualifies() {
			fmt.Println("Second ", i)
			for _, s := range m.M {
				fmt.Println(s)
			}
			fmt.Println()
		}
	}
}
