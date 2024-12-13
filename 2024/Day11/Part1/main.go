package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
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
	scanner.Scan()
	stones := ParseLine(scanner.Text())

	for range 25 {
		for e := stones.Front(); e != nil; e = e.Next() {
			if e.Value == "0" {
				e.Value = "1"
			} else if len(e.Value.(string))%2 == 0 {
				s := e.Value.(string)
				cutPoint := len(s) / 2
				left, right := s[:cutPoint], s[cutPoint:]
				right = strings.TrimLeft(right, "0")
				if right == "" {
					right = "0"
				}
				e.Value = left
				stones.InsertAfter(right, e)
				e = e.Next()
			} else {
				num, err := strconv.Atoi(e.Value.(string))
				if err != nil {
					panic(err)
				}
				num *= 2024
				e.Value = fmt.Sprint(num)
			}
		}
	}
	fmt.Println("Amount of stones:", stones.Len())
}

func ParseLine(line string) *list.List {
	l := list.New()
	for _, s := range strings.Split(line, " ") {
		l.PushBack(s)
	}
	return l
}

func PrintStones(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
