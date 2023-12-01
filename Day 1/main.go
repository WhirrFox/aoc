package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		fir, sec := getNumbers2(scanner.Text())
		sum += fir*10 + sec
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func getNumbers(str string) (int, int) {
	var first, second int
	for _, s := range str {
		if i, err := strconv.Atoi(string(s)); err == nil {
			if first == 0 {
				first = i
			}
			second = i
		}
	}
	return first, second
}

var re = regexp.MustCompile(`(?m)(\d|one|two|three|four|five|six|seven|eight|nine)`)

func getNumbers2(str string) (int, int) {
	var first, last, i int
	for {
		match := re.FindString(str[i:])
		if match == "" {
			return first, last
		}

		if first == 0 {
			first = strToInt(match)
		}

		last = strToInt(match)
		i++
	}
}

func strToInt(str string) int {
	if i, err := strconv.Atoi(string(str)); err == nil {
		return i
	}
	switch str {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	return -1
}
