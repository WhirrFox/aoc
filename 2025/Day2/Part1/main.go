package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sum = 0

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	ranges := strings.Split(scanner.Text(), ",")
	for _, r := range ranges {
		values := strings.Split(r, "-")
		start, _ := strconv.Atoi(values[0])
		end, _ := strconv.Atoi(values[1])
		for i := start; i <= end; i++ {
			checkID(i)
		}
	}

	fmt.Println(sum)
}

func checkID(id int) {
	str := strconv.Itoa(id)
	if len(str)%2 == 1 {
		return
	}

	if str[:len(str)/2] == str[len(str)/2:] {
		sum += id
	}
}
