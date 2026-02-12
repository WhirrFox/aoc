package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

var ingredientsRanges = []IngredientsRange{}

func main() {
	readInput()

	slices.SortFunc(ingredientsRanges, func(a, b IngredientsRange) int {
		return cmp.Compare(int(a.from), int(b.from))
	})

	SquashIngredientsRanges()

	sum := 0
	for _, r := range ingredientsRanges {
		sum += r.Length()
	}
	fmt.Println("Fresh ingredients:", sum)
}

func SquashIngredientsRanges() {
	newIngredientsRanges := []IngredientsRange{}
	mergeMade := true
	for mergeMade {
		mergeMade = false
		for i := 0; i < len(ingredientsRanges); i++ {
			if i == len(ingredientsRanges)-1 {
				newIngredientsRanges = append(newIngredientsRanges, ingredientsRanges[i])
			} else if ingredientsRanges[i].IsOverlapping(&ingredientsRanges[i+1]) {
				newIngredientsRanges = append(newIngredientsRanges,
					MergeRange(ingredientsRanges[i], ingredientsRanges[i+1]))
				mergeMade = true
				i++ // to skip the second one
			} else {
				newIngredientsRanges = append(newIngredientsRanges, ingredientsRanges[i])
			}
		}
		ingredientsRanges = newIngredientsRanges
		newIngredientsRanges = []IngredientsRange{}
	}
}

const (
	ReadingIngredientsRanges = iota
	ReadingIngredients
)

func readInput() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanMode := ReadingIngredientsRanges
	for scanner.Scan() {
		switch scanMode {
		case ReadingIngredientsRanges:
			if scanner.Text() == "" {
				scanMode++
				continue
			}
			ingredientsRanges = append(ingredientsRanges, NewIngredientsRange(scanner.Text()))
		case ReadingIngredients:
			return // Not relevant for part 2
		}
	}
}
