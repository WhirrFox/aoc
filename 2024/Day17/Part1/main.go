package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	p := Processor{}
	program := []int{}
	lineNum := 0
	for scanner.Scan() {
		if lineNum == 3 || lineNum > 4 {
			lineNum++
			continue
		}

		value := strings.Split(scanner.Text(), ":")[1]
		if lineNum == 4 {
			for _, s := range strings.Split(strings.Trim(value, " "), ",") {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				program = append(program, n)
			}
		} else {
			n, err := strconv.Atoi(value[1:])
			if err != nil {
				panic(err)
			}
			switch lineNum {
			case 0:
				p.RegisterA = n
			case 1:
				p.RegisterB = n
			case 2:
				p.RegisterC = n
			}
		}

		lineNum++
	}

	p.RunProgram(program)
	fmt.Println(p.GetOutput())
}
