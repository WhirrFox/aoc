package main

import (
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

func (r *IngredientsRange) Length() int {
	return int(r.to) - int(r.from) + 1
}

func (r *IngredientsRange) InRange(i Ingredient) bool {
	return r.from <= i && i <= r.to
}

func (r *IngredientsRange) IsOverlapping(r2 *IngredientsRange) bool {
	return r.InRange(r2.from) || r.InRange(r2.to)
}

func MergeRange(r1, r2 IngredientsRange) IngredientsRange {
	return IngredientsRange{
		MinIngredient(r1.from, r2.from),
		MaxIngredient(r1.to, r2.to)}
}

func MaxIngredient(i1, i2 Ingredient) Ingredient {
	if i1 > i2 {
		return i1
	}
	return i2
}

func MinIngredient(i1, i2 Ingredient) Ingredient {
	if i1 < i2 {
		return i1
	}
	return i2
}
