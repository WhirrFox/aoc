package main

import (
	"strconv"
	"strings"
)

type XY struct {
	X int
	Y int
}

type Robot struct {
	Pos      XY
	Velocity XY
}

func NewRobot(line string) Robot {
	split := strings.Split(line, " ")
	pos := strings.Split(split[0][2:], ",")
	vel := strings.Split(split[1][2:], ",")
	px, err := strconv.Atoi(pos[0])
	if err != nil {
		panic(err)
	}
	py, err := strconv.Atoi(pos[1])
	if err != nil {
		panic(err)
	}
	vx, err := strconv.Atoi(vel[0])
	if err != nil {
		panic(err)
	}
	vy, err := strconv.Atoi(vel[1])
	if err != nil {
		panic(err)
	}
	return Robot{
		Pos: XY{
			X: px,
			Y: py,
		},
		Velocity: XY{
			X: vx,
			Y: vy,
		},
	}
}

func (r *Robot) Move(m *MapMap, seconds int) {
	for range seconds {
		r.Pos.X += r.Velocity.X
		for r.Pos.X < 0 {
			r.Pos.X += m.Width
		}
		for r.Pos.X >= m.Width {
			r.Pos.X -= m.Width
		}

		r.Pos.Y += r.Velocity.Y
		for r.Pos.Y < 0 {
			r.Pos.Y += m.Height
		}
		for r.Pos.Y >= m.Height {
			r.Pos.Y -= m.Height
		}
	}
}

func (r *Robot) GetQuadrant(m *MapMap) int {
	midWidth := m.Width / 2
	midHeight := m.Height / 2
	if r.Pos.X < midWidth && r.Pos.Y < midHeight {
		return 1
	}
	if r.Pos.X > midWidth && r.Pos.Y < midHeight {
		return 2
	}
	if r.Pos.X < midWidth && r.Pos.Y > midHeight {
		return 3
	}
	if r.Pos.X > midWidth && r.Pos.Y > midHeight {
		return 4
	}
	return -1
}
