package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		sum += parseGame(scanner.Text())
	}

	fmt.Println(sum)
}

func parseGame(str string) int {
	substr := strings.Split(str, ":")
	gameId, err := strconv.Atoi(strings.Split(substr[0], " ")[1])
	if err != nil {
		log.Fatal(err)
	}
	subsets := strings.Split(substr[1], ";")
	for _, s := range subsets {
		if !subsetPossible(s) {
			return 0
		}
	}

	return gameId
}

func subsetPossible(str string) bool {
	for _, cube := range strings.Split(str, ",") {
		if i := strings.Index(cube, "red"); i != -1 {
			if getNum(cube, i) > 12 {
				return false
			}
		} else if i := strings.Index(cube, "green"); i != -1 {
			if getNum(cube, i) > 13 {
				return false
			}
		} else if i := strings.Index(cube, "blue"); i != -1 {
			if getNum(cube, i) > 14 {
				return false
			}
		} else {
			log.Fatal(cube)
		}
	}
	return true
}

func getNum(str string, i int) int {
	num, err := strconv.Atoi(str[1 : i-1])
	if err != nil {
		log.Fatal(err)
	}
	return num
}
