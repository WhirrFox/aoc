package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	patterns := []string{}
	designs := []string{}

	step := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			step++
			continue
		}

		switch step {
		case 0:
			patterns = strings.Split(scanner.Text(), ", ")
		case 1:
			designs = append(designs, scanner.Text())
		}
	}

	amount := 0
	for _, design := range designs {
		if PatternPossible(design, patterns) {
			amount++
		}
	}

	fmt.Println("Possible Designs:", amount)
}

func PatternPossible(design string, patterns []string) bool {
	for i := len(design); i > 0; i-- {
		d := design[:i]
		if slices.Contains(patterns, d) {
			remainDesign := design[len(d):]
			if remainDesign == "" {
				return true
			}
			if PatternPossible(remainDesign, patterns) {
				return true
			}
		}
	}
	return false
}
