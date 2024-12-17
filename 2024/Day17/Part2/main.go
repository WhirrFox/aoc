package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	programString := ""
	lineNum := 0
	for scanner.Scan() {
		if lineNum == 3 || lineNum > 4 {
			lineNum++
			continue
		}

		value := strings.Split(scanner.Text(), ":")[1]
		if lineNum == 4 {
			programString = strings.Trim(value, " ")
			for _, s := range strings.Split(programString, ",") {
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

	a := 0
	revProgram := slices.Clone(program)
	slices.Reverse(revProgram)

	currentIndex := 0
	currentProgram := fmt.Sprint(revProgram[currentIndex])
	for {
		p2 := p
		p2.RegisterA = a
		p2.RunProgram(program)
		if p2.GetOutput() == currentProgram {
			fmt.Printf("%15d   %s\n", a, p2.GetOutput())
			currentIndex++
			if currentIndex >= len(revProgram) {
				return
			}
			currentProgram = fmt.Sprint(revProgram[currentIndex]) + "," + currentProgram
			a = a << 3
		} else {
			a++
		}
	}
}
