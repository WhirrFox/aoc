package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cards = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

type Hand struct {
	Hand string
	Type int
	Bet  int
}

var hands = []Hand{}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		bet, _ := strconv.Atoi(str[1])
		hands = append(hands, Hand{
			Hand: str[0],
			Type: getType(str[0]),
			Bet:  bet,
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type != hands[j].Type {
			return hands[i].Type < hands[j].Type
		}
		for c := 0; c < 5; c++ {
			if hands[i].Hand[c] != hands[j].Hand[c] {
				return cardHigher(string(hands[i].Hand[c]), string(hands[j].Hand[c]))
			}
		}
		log.Fatal("No!")
		return false
	})

	sum := 0
	for i, h := range hands {
		fmt.Println(i+1, h.Type, h.Hand)
		sum += (i + 1) * h.Bet
	}
	fmt.Println(sum, "Onegai~")
}

func getType(hand string) int {
	ofAKind := getOfAKind(hand)
	if ofAKind > 4 {
		return ofAKind + 2
	} else if ofAKind < 4 {
		return ofAKind
	} else {
		if isFullHouse(hand) {
			return 5
		}
		return 4
	}
}

func getOfAKind(hand string) int {
	maxKind := 0
	twoPair := 0
	for _, c := range cards {
		count := strings.Count(hand, c)

		if count > maxKind {
			maxKind = count
		}

		if count == 2 {
			twoPair++
		}
	}
	if twoPair >= 2 {
		return 3
	}
	if maxKind < 3 {
		return maxKind
	}
	return maxKind + 1
}

func isFullHouse(hand string) bool {
	for _, c := range cards {
		if count := strings.Count(hand, c); count == 2 {
			return true
		}
	}
	return false
}

func cardHigher(card1, card2 string) bool {
	for _, c := range cards {
		if card1 == c {
			return false
		} else if card2 == c {
			return true
		}
	}
	log.Fatal("No!")
	return false
}
