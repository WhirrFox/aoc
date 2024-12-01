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
	errHandle(err)
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sl := strToIntSlice(scanner.Text())
		hist := [][]int{sl}
		fillSequences(&hist)
		sum += getNewNum(&hist)
	}

	fmt.Println(sum)
}

func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func strToIntSlice(str string) []int {
	slice := []int{}
	for _, s := range strings.Split(str, " ") {
		i, err := strconv.Atoi(s)
		errHandle(err)
		slice = append(slice, i)
	}
	return slice
}

func fillSequences(slice *[][]int) {
	up, down := 0, len((*slice)[0])
	for {
		sum, newSlice := 0, []int{}
		for i := 0; i < down-1; i++ {
			num := (*slice)[up][i+1] - (*slice)[up][i]
			sum += num
			newSlice = append(newSlice, num)
		}
		*slice = append(*slice, newSlice)

		if sum == 0 {
			return
		}

		up++
		down--
	}
}

func getNewNum(slice *[][]int) int {
	num := 0
	for i := len(*slice) - 2; i >= 0; i-- {
		row := (*slice)[i]
		num = num + row[len(row)-1]
	}
	return num
}
