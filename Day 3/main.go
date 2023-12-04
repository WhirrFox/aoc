package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

const width, height = 140, 140

type schemT [height]string

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	schem := readSchematic(scanner)

	sum := 0
	for i, _ := range schem {
		sum += parseRow(&schem, i)
	}

	fmt.Println("Sum", sum)
}

func readSchematic(scanner *bufio.Scanner) schemT {
	schem := schemT{}
	i := 0
	for scanner.Scan() {
		schem[i] = scanner.Text()
		i++
	}
	return schem
}

var exp = regexp.MustCompile(`(?m)(\d+)`)

func parseRow(schem *schemT, row int) int {
	sum := 0
	match := exp.FindAllIndex([]byte(schem[row]), -1)
	for _, i := range match {
		if checkAdjacent(schem, i[0], row) || checkAdjacent(schem, i[1]-1, row) {
			num, err := strconv.Atoi(schem[row][i[0]:i[1]])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(num)
			sum += num
		}
	}
	return sum
}

func checkAdjacent(schem *schemT, col, row int) bool {
	isPartNum := 0
	if row > 0 {
		if col > 0 {
			// Top left
			isPartNum += isSymbol(schem, col-1, row-1)
		}
		// Top middle
		isPartNum += isSymbol(schem, col, row-1)

		if col < width-1 {
			// Top right
			isPartNum += isSymbol(schem, col+1, row-1)
		}
	}

	if col > 0 {
		// Middle left
		isPartNum += isSymbol(schem, col-1, row)
	}
	if col < width-1 {
		// Middle right
		isPartNum += isSymbol(schem, col+1, row)
	}

	if row < height-1 {
		if col > 0 {
			// Bottom left
			isPartNum += isSymbol(schem, col-1, row+1)
		}
		// Bottom middle
		isPartNum += isSymbol(schem, col, row+1)

		if col < width-1 {
			// Bottom right
			isPartNum += isSymbol(schem, col+1, row+1)
		}
	}

	return isPartNum > 0
}

func isSymbol(schem *schemT, col, row int) int {
	if schem[row][col] != '.' && !unicode.IsDigit(rune(schem[row][col])) {
		return 1
	}
	return 0
}
