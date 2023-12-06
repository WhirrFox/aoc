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

func main() {
	file, err := os.Open("input.txt")
	errHandle(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	td := readInput(scanner)

	fmt.Println(calc(td))
}

func readInput(scanner *bufio.Scanner) TimeDist {
	scanner.Scan()
	time := strings.Split(scanner.Text(), ":")[1]
	scanner.Scan()
	dist := strings.Split(scanner.Text(), ":")[1]

	time = strings.ReplaceAll(time, " ", "")
	dist = strings.ReplaceAll(dist, " ", "")

	t, err := strconv.Atoi(time)
	errHandle(err)
	d, err := strconv.Atoi(dist)
	errHandle(err)

	return TimeDist{
		Time: t,
		Dist: d,
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
