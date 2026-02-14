package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '+' || line[0] == '*' {
			SetOperations(line)
		} else {
			AddNumbers(line)
		}
	}

	total := 0
	for _, p := range problems {
		total += p.Solve()
	}
	fmt.Println("Grand total:", total)
}

func FilterEmptyOut(in []string) (out []string) {
	out = in[:0]
	for _, s := range in {
		if s != "" {
			out = append(out, s)
		}
	}
	return
}

func AddNumbers(line string) {
	nums := FilterEmptyOut(strings.Split(line, " "))
	if len(problems) == 0 {
		problems = make([]Problem, len(nums))
	}
	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		problems[i].AddNum(n)
	}
}

func SetOperations(line string) {
	operations := FilterEmptyOut(strings.Split(line, " "))
	for i, o := range operations {
		problems[i].Operation = Operation(o)
	}
}
