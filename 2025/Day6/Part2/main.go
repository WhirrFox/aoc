package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Operation string

const (
	ADDITION       = "+"
	MULTIPLICATION = "*"
)

type Problem struct {
	Numbers   []int
	Operation Operation
}

func (p *Problem) AddNum(i int) {
	p.Numbers = append(p.Numbers, i)
}

func (p *Problem) Solve() (solution int) {
	switch p.Operation {
	case ADDITION:
		solution = 0
		for _, n := range p.Numbers {
			solution += n
		}
	case MULTIPLICATION:
		solution = 1
		for _, n := range p.Numbers {
			solution *= n
		}
	default:
		panic("Operation not supported")
	}
	return
}

var problems = []Problem{}
var grid = Grid{}

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	operationsLine := len(grid) - 1
	indexes := GetOperationsIndexes()
	for index, in := range indexes {
		start := in[0]
		var end int
		if index == len(indexes)-1 {
			end = start + 10
		} else {
			end = indexes[index+1][0]
		}

		problems = append(problems, ParseProblem(start, end, operationsLine))
	}

	total := 0
	for _, p := range problems {
		total += p.Solve()
	}
	fmt.Println("Grand total:", total)
}

func ParseProblem(start, end, operationsLine int) (p Problem) {
	p.Operation = Operation(string(grid.At(start, operationsLine)))
	for x := start; x < end; x++ {
		num := []rune{}
		for y := range operationsLine {
			if r := grid.At(x, y); r != ' ' {
				num = append(num, r)
			}
		}
		if len(num) == 0 {
			return
		}

		i, err := strconv.Atoi(string(num))
		if err != nil {
			panic(err)
		}

		p.AddNum(i)
	}
	return
}

func GetOperationsIndexes() [][]int {
	re := regexp.MustCompile(`(?m)[*+]`)
	line := string(grid[len(grid)-1])
	return re.FindAllStringIndex(line, -1)
}
