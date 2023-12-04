package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const maxCards = 190

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cardCounts := [maxCards]*int{}
	for i := 0; i < maxCards; i++ {
		val := 1
		cardCounts[i] = &val
	}

	scanner := bufio.NewScanner(file)
	sum, card := 0, 0
	for scanner.Scan() {
		for i := 0; i < *cardCounts[card]; i++ {
			processCard(&cardCounts, card, scanner.Text())
			sum++
		}
		card++
	}

	fmt.Println("Sum", sum)
}

func processCard(cardCounts *[maxCards]*int, card int, s string) {
	fmt.Println("Card", card+1)
	sum := 0
	s = strings.Split(s, ":")[1]
	split := strings.Split(s, "|")
	winNum := strings.Split(split[0], " ")
	num := strings.Split(split[1], " ")

	for _, n := range num {
		if n == "" {
			continue
		}
		if slices.Contains(winNum, n) {
			sum++
		}
	}

	for c := card + 1; c <= card+sum; c++ {
		if c >= 190 {
			break
		}
		*cardCounts[c]++
	}
}
