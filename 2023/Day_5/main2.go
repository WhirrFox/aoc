package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type DestSrc struct {
	Dest int
	Src  int
	Len  int
}

var maps = [7][]DestSrc{}
var seeds = []int{}

func main() {
	file, err := os.Open("input.txt")
	errHandle(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	readSeeds(scanner)
	scanner.Scan()

	step := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			step++
			continue
		} else if strings.Contains(scanner.Text(), "-") {
			continue
		}

		parseMapRow(step, scanner.Text())
	}

	var wg sync.WaitGroup
	first, lowestAll := -1, math.MaxInt
	for _, s := range seeds {
		if first == -1 {
			first = s
			continue
		}

		wg.Add(1)
		go func(first, second int) {
			defer wg.Done()

			lowest := math.MaxInt
			for i := first; i < first+second; i++ {
				num := getNumber(6, i)
				if num < lowest {
					lowest = num
				}
			}

			fmt.Println(lowest)
			if lowest < lowestAll {
				lowestAll = lowest
			}
		}(first, s)
		first = -1
	}
	wg.Wait()

	fmt.Println("Result", lowestAll)
}

func readSeeds(scanner *bufio.Scanner) {
	str := strings.Split(scanner.Text(), ":")[1]
	for _, s := range strings.Split(str, " ") {
		if s == "" {
			continue
		}

		i, err := strconv.Atoi(s)
		errHandle(err)
		seeds = append(seeds, i)
	}
}

func parseMapRow(step int, row string) {
	r := strings.Split(row, " ")
	dest, err := strconv.Atoi(r[0])
	errHandle(err)
	src, err := strconv.Atoi(r[1])
	errHandle(err)
	len, err := strconv.Atoi(r[2])
	errHandle(err)

	maps[step] = append(maps[step], DestSrc{
		Dest: dest,
		Src:  src,
		Len:  len,
	})
}

func getNumber(step, num int) int {
	if step > 0 {
		num = getNumber(step-1, num)
	}
	for _, m := range maps[step] {
		if num >= m.Src && num < m.Src+m.Len {
			return mapInt(num, m.Src, m.Src+m.Len, m.Dest, m.Dest+m.Len)
		}
	}

	return num
}

func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func mapInt(value, low1, high1, low2, high2 int) int {
	return low2 + (value-low1)*(high2-low2)/(high1-low1)
}
