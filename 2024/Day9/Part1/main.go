package main

import (
	"bufio"
	"fmt"
	"os"
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
	diskMap := ParseLine(scanner.Text())
	diskMap = Defrag(diskMap)
	checksum := CalcChecksum(diskMap)
	fmt.Println("Checksum:", checksum)
}

func ParseLine(line string) (output []int) {
	id := 0
	isBlock := true
	for _, r := range line {
		length := int(r - '0')
		if isBlock {
			for range length {
				output = append(output, id)
			}
			id++
		} else {
			for range length {
				output = append(output, -1)
			}
		}
		isBlock = !isBlock
	}
	return
}

func Defrag(diskMap []int) []int {
	left, right := 0, len(diskMap)-1
	for {
		for diskMap[left] != -1 {
			left++
			if left > right {
				return diskMap
			}
		}
		for diskMap[right] == -1 {
			right--
			if left > right {
				return diskMap
			}
		}
		diskMap[left], diskMap[right] = diskMap[right], diskMap[left]
	}
}

func CalcChecksum(diskMap []int) (sum int) {
	for i, n := range diskMap {
		if n == -1 {
			return
		}
		sum += i * n
	}
	panic("This shouldn't happen")
}
