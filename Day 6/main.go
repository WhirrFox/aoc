package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type TimeDist struct {
	Time int
	Dist int
}

var doc = []TimeDist{}

func main() {
	file, err := os.Open("input.txt")
	errHandle(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	readInput(scanner)
	pro := 1
	for _, d := range doc {
		pro *= calc(d)
	}

	fmt.Println(pro)
}

func readInput(scanner *bufio.Scanner) {
	scanner.Scan()
	time := strings.Split(strings.Split(scanner.Text(), ":")[1], " ")
	scanner.Scan()
	dist := strings.Split(strings.Split(scanner.Text(), ":")[1], " ")
	for _, s := range time {
		if s == "" {
			continue
		}

		t, err := strconv.Atoi(s)
		errHandle(err)

		doc = append(doc, TimeDist{
			Time: t,
		})
	}

	i := 0
	for _, s := range dist {
		if s == "" {
			continue
		}

		d, err := strconv.Atoi(s)
		errHandle(err)

		doc[i].Dist = d
		i++
	}
}

func calc(td TimeDist) int {
	count := 0
	for i := 1; i < td.Time; i++ {
		dist := (td.Time - i) * i
		if dist > td.Dist {
			count++
		}
	}
	return count
}

func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
