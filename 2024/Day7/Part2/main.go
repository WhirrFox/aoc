package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Equation struct {
	TestValue int
	Numbers   []int
}

const (
	ADD = iota + 0
	MULTIPLY
	CONCAT
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

	threads := 4
	wg := sync.WaitGroup{}
	chunkSize := (len(equations) + threads - 1) / threads
	for i := 0; i < len(equations); i += chunkSize {
		end := i + chunkSize
		if end > len(equations) {
			end = len(equations)
		}
		wg.Add(1)
		go func(eqs []Equation) {
			localTotal := 0
			for _, e := range eqs {
				if IsTestValueCorrect(e) {
					localTotal += e.TestValue
				}
			}

			total += localTotal
			wg.Done()
		}(equations[i:end])
	}

	wg.Wait()
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

			if op%3 == ADD {
				result += n
			} else if op%3 == MULTIPLY {
				result *= n
			} else {
				result, _ = strconv.Atoi(fmt.Sprint(result, "", n))
			}

			op /= 3
		}

		if result == e.TestValue {
			return true
		}

		operation++
		if operation > 1000000 {
			return false
		}
	}
}
