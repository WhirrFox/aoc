package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	Lights  []bool
	Buttons [][]int
}

func NewMachine(line string) (m Machine) {
	fields := strings.Split(line, " ")
	for _, b := range fields[0][1 : len(fields[0])-1] {
		switch b {
		case '.':
			m.Lights = append(m.Lights, false)
		case '#':
			m.Lights = append(m.Lights, true)
		default:
			panic("Character not supported")
		}
	}

	for _, buttonList := range fields[1 : len(fields)-1] {
		buttons := []int{}
		buttonStrings := strings.Split(buttonList[1:len(buttonList)-1], ",")
		for _, b := range buttonStrings {
			i, err := strconv.Atoi(b)
			if err != nil {
				panic(err)
			}
			buttons = append(buttons, i)
		}
		m.Buttons = append(m.Buttons, buttons)
	}
	return
}

func (m *Machine) CheckMachine() int {
	for i := range 100 {
		if m.CheckAllButtons(i, []int{}) {
			return i
		}
	}
	panic("Infinite loop with machine")
}

func (m *Machine) CheckAllButtons(length int, combinations []int) bool {
	if len(combinations) == length {
		return m.TurnsOn(combinations)
	}

	combinations = append(combinations, -1)
	for i := range m.Buttons {
		combinations[len(combinations)-1] = i
		if m.CheckAllButtons(length, combinations) {
			return true
		}
	}
	return false
}

func (m *Machine) TurnsOn(combinations []int) bool {
	lights := make([]bool, len(m.Lights))
	for _, c := range combinations {
		buttons := m.Buttons[c]
		for _, b := range buttons {
			lights[b] = !lights[b]
		}
	}

	return slices.Equal(m.Lights, lights)
}

var machines = []Machine{}

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		machines = append(machines, NewMachine(scanner.Text()))
	}

	buttonPresses := 0
	for _, m := range machines {
		buttonPresses += m.CheckMachine()
	}
	fmt.Println("Required button presses:", buttonPresses)
}
