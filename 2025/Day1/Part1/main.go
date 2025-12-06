package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	dial := 50
	password := 0
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
			dial -= num
		case 'R':
			dial += num
		}

		tempDial := dial
		if math.Abs(float64(tempDial)) >= 100 {
			leftHalf := tempDial / 100
			tempDial -= leftHalf * 100
		}
		if tempDial == 0 {
			password++
		}

		// fmt.Println(text, dial, tempDial)
	}
	fmt.Println("Password:", password)
}
