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

	patterns := make(map[int][]string)
	designs := []string{}

	step := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			step++
			continue
		}

		switch step {
		case 0:
			rawPatterns := strings.Split(scanner.Text(), ", ")
			slices.Sort(rawPatterns)
			for _, p := range rawPatterns {
				if _, exists := patterns[len(p)]; !exists {
					patterns[len(p)] = []string{}
				}
				patterns[len(p)] = append(patterns[len(p)], p)
			}
		case 1:
			designs = append(designs, scanner.Text())
		}
	}

	amount := 0
	for _, design := range designs {
		possiblePatterns := 0
		possiblePatterns += PossibleWays(design, &patterns)
		amount += possiblePatterns
	}

	fmt.Println("All different ways:", amount)
}

var possibleWaysCache = make(map[string]int)

func PossibleWays(design string, patterns *map[int][]string) (possiblePatterns int) {
	if val, exists := possibleWaysCache[design]; exists {
		return val
	}

	for i := min(8, len(design)); i > 0; i-- {
		d := design[:i]

		if _, found := slices.BinarySearch((*patterns)[len(d)], d); found {
			remainDesign := design[len(d):]
			if remainDesign == "" {
				possiblePatterns++
			} else {
				possiblePatterns += PossibleWays(remainDesign, patterns)
			}
		}
	}
	possibleWaysCache[design] = possiblePatterns
	return possiblePatterns
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
