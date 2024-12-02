package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report []int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	reports := []Report{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		reports = append(reports, parseLine(scanner.Text()))
	}

	safeReports := 0
	for _, report := range reports {
		if isSafeReport(report) {
			safeReports++
		}
	}
	fmt.Println("Safe Reports: ", safeReports)
}

func parseLine(text string) Report {
	stringSlice := strings.Split(text, " ")
	report := make(Report, len(stringSlice))
	for i, s := range stringSlice {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		report[i] = n
	}
	return report
}

type Mode int

const (
	INCREASING = iota + 1
	DECREASING
)

func isSafeReport(report Report) bool {
	mode := findMode(report)

	for i, n := range report {
		if i == len(report)-1 {
			return true
		}

		nextN := report[i+1]
		diff := abs(n - nextN)

		if !(diff >= 1 && diff <= 3) {
			return false
		}

		switch mode {
		case INCREASING:
			if !(n < nextN) {
				return false
			}
		case DECREASING:
			if !(n > nextN) {
				return false
			}
		}
	}
	panic("How did we get here?")
}

func findMode(report Report) Mode {
	for i, n := range report {
		if n == report[i+1] {
			continue
		}

		if n < report[i+1] {
			return INCREASING
		} else {
			return DECREASING
		}
	}
	panic("Mode not found")
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
