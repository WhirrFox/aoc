package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const width, height = 140, 140

var sum int

type schemT [height]string
type star struct {
	col       int
	row       int
	firstNum  int
	secondNum int
}

var allStars []*star = []*star{}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	schem := readSchematic(scanner)

	for i, _ := range schem {
		parseRow(&schem, i)
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

func parseRow(schem *schemT, row int) {
	match := exp.FindAllIndex([]byte(schem[row]), -1)
	for _, i := range match {
		currentNum, _ := strconv.Atoi(schem[row][i[0]:i[1]])
		checkAdjacent(schem, i[0], row, currentNum)
		checkAdjacent(schem, i[1]-1, row, currentNum)
	}
}

func checkAdjacent(schem *schemT, col, row, currentNum int) {
	if row > 0 {
		if col > 0 {
			// Top left
			handleStar(schem, col-1, row-1, currentNum)
		}
		// Top middle
		handleStar(schem, col, row-1, currentNum)

		if col < width-1 {
			// Top right
			handleStar(schem, col+1, row-1, currentNum)
		}
	}

	if col > 0 {
		// Middle left
		handleStar(schem, col-1, row, currentNum)
	}
	if col < width-1 {
		// Middle right
		handleStar(schem, col+1, row, currentNum)
	}

	if row < height-1 {
		if col > 0 {
			// Bottom left
			handleStar(schem, col-1, row+1, currentNum)
		}
		// Bottom middle
		handleStar(schem, col, row+1, currentNum)

		if col < width-1 {
			// Bottom right
			handleStar(schem, col+1, row+1, currentNum)
		}
	}
}

func handleStar(schem *schemT, col, row, currentNum int) {
	if schem[row][col] != '*' {
		return
	}

	for _, ls := range allStars {
		if ls.col == col && ls.row == row {
			if ls.secondNum == 0 && ls.firstNum != currentNum {
				sum += ls.firstNum * currentNum
				ls.secondNum = currentNum
				fmt.Println(ls.firstNum, ls.secondNum)
			}
			return
		}
	}

	allStars = append(allStars, &star{
		col:      col,
		row:      row,
		firstNum: currentNum,
	})
}
