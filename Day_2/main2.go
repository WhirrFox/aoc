package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cubes struct {
	Red   int
	Green int
	Blue  int
}

func (c *Cubes) MinCubes(c2 Cubes) {
	c.Red = int(math.Max(float64(c.Red), float64(c2.Red)))
	c.Green = int(math.Max(float64(c.Green), float64(c2.Green)))
	c.Blue = int(math.Max(float64(c.Blue), float64(c2.Blue)))
}

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
	subsets := strings.Split(substr[1], ";")

	c := Cubes{}
	for _, s := range subsets {
		c.MinCubes(parseSubset(s))
	}

	return c.Red * c.Blue * c.Green
}

func parseSubset(str string) Cubes {
	subset := Cubes{}
	for _, cube := range strings.Split(str, ",") {
		if i := strings.Index(cube, "red"); i != -1 {
			subset.Red = getNum(cube, i)
		} else if i := strings.Index(cube, "green"); i != -1 {
			subset.Green = getNum(cube, i)
		} else if i := strings.Index(cube, "blue"); i != -1 {
			subset.Blue = getNum(cube, i)
		} else {
			log.Fatal(cube)
		}
	}

	return subset
}

func getNum(str string, i int) int {
	num, err := strconv.Atoi(str[1 : i-1])
	if err != nil {
		log.Fatal(err)
	}
	return num
}
