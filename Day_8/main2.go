package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var nav string
var m = map[string][2]string{}
var pointers = []*string{}

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

		if strings.HasSuffix(name, "A") {
			pointers = append(pointers, &name)
		}
	}

	for _, p := range pointers {
		step, i := 0, 0
		for {
			if i >= len(nav) {
				i = 0
			}
			switch nav[i] {
			case 'L':
				*p = m[*p][0]
			case 'R':
				*p = m[*p][1]
			default:
				log.Fatal(nav[i], "not found")
			}

			step++
			i++

			if strings.HasSuffix(*p, "Z") {
				fmt.Println(step)
				break
			}
		}

	}

	fmt.Println("Put these numbers in a LCM Calculator online, and the result is the answer!")
	fmt.Println("(I was too lazy to implement a LCM Calculator here)")
}
