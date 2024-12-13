package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type Machine struct {
	ButtonA XY
	ButtonB XY
	Prize   XY
}

type XY struct {
	X int
	Y int
}

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	machines := []Machine{}
	step := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		coordsLine := strings.Split(scanner.Text(), ":")[1]
		coords := strings.Split(coordsLine, ",")
		x, err := strconv.Atoi(strings.Trim(coords[0], " ")[2:])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strings.Trim(coords[1], " ")[2:])
		if err != nil {
			panic(err)
		}
		switch step {
		case 0:
			machines = append(machines, Machine{
				ButtonA: XY{
					X: x,
					Y: y,
				},
			})
		case 1:
			machines[len(machines)-1].ButtonB = XY{
				X: x,
				Y: y,
			}
		case 2:
			machines[len(machines)-1].Prize = XY{
				X: x,
				Y: y,
			}
		}

		step++
		if step > 2 {
			step = 0
		}
	}

	sum := 0
	for _, m := range machines {
		if s := SolveMachine(m); s != -1 {
			sum += s
		}
	}
	fmt.Println("Total Tokens:", sum)
}

func SolveMachine(m Machine) int {
	A := mat.NewDense(2, 2, []float64{
		float64(m.ButtonA.X), float64(m.ButtonB.X),
		float64(m.ButtonA.Y), float64(m.ButtonB.Y),
	})

	B := mat.NewVecDense(2, []float64{
		float64(m.Prize.X),
		float64(m.Prize.Y),
	})

	var x mat.VecDense

	err := x.SolveVec(A, B)
	if err != nil {
		panic(err)
	}

	a := x.At(0, 0)
	b := x.At(1, 0)

	if isWholeNumber(a) && isWholeNumber(b) {
		// I hate float errors and numbers like 47.99999999999999, which math.Trunc() converts to 47 instead of 48
		return int(3*math.Round(a) + math.Round(b))
	}
	return -1
}

func isWholeNumber(a float64) bool {
	const epsilon = 1e-9 // Margin of error
	_, frac := math.Modf(math.Abs(a))
	return frac < epsilon || frac > 1.0-epsilon
}
