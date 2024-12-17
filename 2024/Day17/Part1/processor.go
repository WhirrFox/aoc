package main

import (
	"fmt"
	"math"
	"strings"
)

type Operand int

const (
	adv = iota + 0 // division, write to A
	bxl            // bitwise XOR of B and literal operand
	bst            // modulo 8, write to B
	jnz            // jump if not zero
	bxc            // bitwose XOR of B and C
	out            // modulo 8, output
	bdv            // division, write to B
	cdv            // division, write to C
)

type Processor struct {
	RegisterA int
	RegisterB int
	RegisterC int
	output    []string
}

func (p *Processor) RunProgram(program []int) {
	inP := 0
	for {
		if inP >= len(program) {
			return
		}

		if r := p.execute(Operand(program[inP]), program[inP+1]); r == -1 {
			inP += 2
		} else {
			inP = r
		}
	}
}

func (p *Processor) execute(op Operand, input int) int {
	switch op {
	case adv:
		p.RegisterA = p.RegisterA / int(math.Pow(2, float64(p.getCombo(input))))
		return -1
	case bxl:
		p.RegisterB = p.RegisterB ^ input
		return -1
	case bst:
		p.RegisterB = p.getCombo(input) % 8
		return -1
	case jnz:
		if p.RegisterA == 0 {
			return -1
		}
		return input
	case bxc:
		p.RegisterB = p.RegisterB ^ p.RegisterC
		return -1
	case out:
		p.output = append(p.output, fmt.Sprint((p.getCombo(input) % 8)))
		return -1
	case bdv:
		p.RegisterB = p.RegisterA / int(math.Pow(2, float64(p.getCombo(input))))
		return -1
	case cdv:
		p.RegisterC = p.RegisterA / int(math.Pow(2, float64(p.getCombo(input))))
		return -1
	}
	panic("Invalid operand")
}

func (p *Processor) getCombo(input int) int {
	if input <= 3 {
		return input
	}
	switch input {
	case 4:
		return p.RegisterA
	case 5:
		return p.RegisterB
	case 6:
		return p.RegisterC
	}
	panic("Not valid input")
}

func (p *Processor) GetOutput() string {
	return strings.Join(p.output, ",")
}
