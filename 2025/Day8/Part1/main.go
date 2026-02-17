package main

import (
	"bufio"
	"cmp"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type JunctionBox struct {
	Id      int
	X, Y, Z int
	Circuit int
}

var junctionBoxCount = 0
var circuitCount = 0
var junctionBoxes = []*JunctionBox{}

func NewJunctionBox(s string) (j JunctionBox) {
	vals := strings.Split(s, ",")
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		panic(err)
	}
	z, err := strconv.Atoi(vals[2])
	if err != nil {
		panic(err)
	}
	j = JunctionBox{
		Id: junctionBoxCount,
		X:  x, Y: y, Z: z,
		Circuit: -1,
	}
	junctionBoxes = append(junctionBoxes, &j)
	junctionBoxCount++
	return
}

func (j *JunctionBox) DistanceTo(j2 *JunctionBox) float64 {
	x := j.X - j2.X
	y := j.Y - j2.Y
	z := j.Z - j2.Z
	return math.Sqrt(float64(x*x + y*y + z*z))
}

type JunctionBoxMatch struct {
	Box1, Box2 *JunctionBox
	Distance   float64
}

type GroupCount struct {
	Circuit, Count int
}

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		NewJunctionBox(scanner.Text())
	}

	ConnectCircuits()
	groups := GroupCircuits()

	j, _ := json.MarshalIndent(groups, "", "  ")
	fmt.Println(string(j))
	fmt.Println("Answer:", groups[len(groups)-2].Count*groups[len(groups)-3].Count*groups[len(groups)-4].Count)
}

func GroupCircuits() []GroupCount {
	boxGroups := make(map[int][]*JunctionBox)
	for _, j := range junctionBoxes {
		boxGroups[j.Circuit] = append(boxGroups[j.Circuit], j)
	}

	groups := []GroupCount{}
	for k, v := range boxGroups {
		groups = append(groups, GroupCount{k, len(v)})
	}
	slices.SortFunc(groups, func(a, b GroupCount) int {
		return cmp.Compare(a.Count, b.Count)
	})
	return groups
}

func ConnectCircuits() {
	circuitsConnected := 0
	for _, m := range GetMatches() {
		circuitsConnected++
		if m.Box1.Circuit == m.Box2.Circuit && m.Box1.Circuit != -1 {
			continue
		}
		if m.Box1.Circuit == -1 && m.Box2.Circuit == -1 {
			m.Box1.Circuit = circuitCount
			m.Box2.Circuit = circuitCount
			circuitCount++
		} else if m.Box1.Circuit != -1 && m.Box2.Circuit == -1 {
			m.Box2.Circuit = m.Box1.Circuit
		} else if m.Box1.Circuit == -1 && m.Box2.Circuit != -1 {
			m.Box1.Circuit = m.Box2.Circuit
		} else if m.Box1.Circuit != m.Box2.Circuit {
			ReplaceCircuits(m.Box2.Circuit, m.Box1.Circuit)
		} else {
			panic("This shouldn't happen?")
		}

		if circuitsConnected == len(junctionBoxes) {
			return
		}
	}
}

func ReplaceCircuits(old, new int) {
	for _, j := range junctionBoxes {
		if j.Circuit == old {
			j.Circuit = new
		}
	}
}

func GetMatches() (matches []JunctionBoxMatch) {
	for i1 := range len(junctionBoxes) {
		for i2 := i1 + 1; i2 < len(junctionBoxes); i2++ {
			d := junctionBoxes[i1].DistanceTo(junctionBoxes[i2])
			matches = append(matches, JunctionBoxMatch{junctionBoxes[i1], junctionBoxes[i2], d})
		}
	}
	slices.SortFunc(matches, func(a, b JunctionBoxMatch) int {
		return cmp.Compare(a.Distance, b.Distance)
	})
	return
}
