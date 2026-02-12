package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ingredient int

type IngredientsRange struct {
	from, to Ingredient
}

func NewIngredientsRange(s string) IngredientsRange {
	inputs := strings.Split(s, "-")
	from, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}
	to, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}
	return IngredientsRange{Ingredient(from), Ingredient(to)}
}

func (r *IngredientsRange) InRange(i Ingredient) bool {
	return r.from <= i && i <= r.to
}

var ingredients = []Ingredient{}
var ingredientsRanges = []IngredientsRange{}

func main() {
	readInput()

	freshIngredients := []Ingredient{}
	for _, i := range ingredients {
		for _, r := range ingredientsRanges {
			if r.InRange(i) {
				freshIngredients = append(freshIngredients, i)
				break
			}
		}
	}

	fmt.Println("Fresh ingredient IDs:", len(freshIngredients))
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
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			ingredients = append(ingredients, Ingredient(num))
		}
	}
}
