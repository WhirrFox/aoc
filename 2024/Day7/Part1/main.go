package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	TestValue int
	Numbers   []int
}

const (
	ADD = iota + 0
	MULTIPLY
)

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	equations := ParseInput(file)

	total := 0
	for _, e := range equations {
		if IsTestValueCorrect(e) {
			total += e.TestValue
		}
	}

	fmt.Println("Total:", total)
}

func ParseInput(r io.Reader) []Equation {
	equations := []Equation{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		values := strings.Split(scanner.Text(), ":")
		testValue, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}

		numbersString := strings.Split(values[1][1:], " ")
		numbers := make([]int, len(numbersString))
		for i, s := range numbersString {
			numbers[i], err = strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
		}

		equations = append(equations, Equation{
			TestValue: testValue,
			Numbers:   numbers,
		})
	}
	return equations
}

func IsTestValueCorrect(e Equation) bool {
	operation := 0
	for {
		result := e.Numbers[0]
		op := operation
		for i, n := range e.Numbers {
			if i == 0 {
				continue
			}

			if op%2 == ADD {
				result += n
			} else {
				result *= n
			}

			op /= 2
		}

		if result == e.TestValue {
			return true
		}

		operation++
		if operation > 10000 {
			return false
		}
	}
}
