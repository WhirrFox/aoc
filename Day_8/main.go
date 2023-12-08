package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var nav string
var m = map[string][2]string{}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	nav = scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		name := line[0:3]
		left := line[7:10]
		right := line[12:15]
		m[name] = [2]string{left, right}
	}

	step, i := 0, 0
	p := "AAA"
	for {
		fmt.Println(p)
		if i >= len(nav) {
			i = 0
		}
		switch nav[i] {
		case 'L':
			p = m[p][0]
		case 'R':
			p = m[p][1]
		default:
			log.Fatal(nav[i], "not found")
		}

		step++
		i++

		if p == "ZZZ" {
			break
		}
	}

	fmt.Println("You look so good~", step)
}
