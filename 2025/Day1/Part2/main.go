package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dial = 50
var password = 0

// I got too lazy figuring out the math, so I just increment through every number lol
func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		num, err := strconv.Atoi(text[1:])
		if err != nil {
			panic(err)
		}
		switch text[0] {
		case 'L':
			goLeft(num)
		case 'R':
			goRight(num)
		}
	}
	fmt.Println("Password:", password)
}

func goLeft(num int) {
	for range num {
		dial--
		switch dial {
		case 0, 100:
			password++
			dial = 0
		case -1:
			dial = 99
		}
	}
}

func goRight(num int) {
	for range num {
		dial++
		switch dial {
		case 0, 100:
			password++
			dial = 0
		case -1:
			dial = 99
		}
	}
}
