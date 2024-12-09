package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Block struct {
	Id     int
	Length int
}

func (b *Block) IsEmptySpace() bool {
	return b.Id == -1
}

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

func ParseLine(line string) (output []Block) {
	id := 0
	isBlock := true
	for _, r := range line {
		length := int(r - '0')
		if isBlock {
			output = append(output, Block{
				Id:     id,
				Length: length,
			})
			id++
		} else {
			output = append(output, Block{
				Id:     -1,
				Length: length,
			})
		}
		isBlock = !isBlock
	}
	return
}

func Defrag(diskMap []Block) []Block {
	id := diskMap[len(diskMap)-1].Id
	if id == -1 {
		panic("Nope")
	}
	for {
		if id == 0 {
			return diskMap
		}
		blockPos := GetBlockPosition(diskMap, id)
		for i, b := range diskMap {
			if !b.IsEmptySpace() {
				continue
			}
			if diskMap[blockPos].Length <= b.Length {
				if i > blockPos {
					break
				}
				remaining := b.Length - diskMap[blockPos].Length
				diskMap[i].Length = diskMap[blockPos].Length
				diskMap[i], diskMap[blockPos] = diskMap[blockPos], diskMap[i]
				if remaining > 0 {
					diskMap = slices.Insert(diskMap, i+1, Block{
						Id:     -1,
						Length: remaining,
					})
					diskMap = CleanUpEmptySpace(diskMap)
				}
				break
			}
		}

		id--
	}
}

func GetBlockPosition(diskMap []Block, id int) int {
	for i, b := range diskMap {
		if b.Id == id {
			return i
		}
	}
	panic("Block not found")
}

func CalcChecksum(diskMap []Block) (sum int) {
	i := 0
	for _, b := range diskMap {
		for range b.Length {
			if !b.IsEmptySpace() {
				sum += i * b.Id
			}
			i++
		}
	}
	return
}

func CleanUpEmptySpace(diskMap []Block) []Block {
	len := len(diskMap) - 1
	for i := 0; i < len; i++ {
		if diskMap[i].IsEmptySpace() && diskMap[i+1].IsEmptySpace() {
			diskMap[i].Length += diskMap[i+1].Length
			diskMap = removeIndex(diskMap, i+1)
			i--
			len--
		}
	}
	return diskMap
}

func removeIndex(s []Block, index int) []Block {
	return append(s[:index], s[index+1:]...)
}
