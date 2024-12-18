package main

type Direction int

const (
	WEST = iota + 0
	SOUTH
	EAST
	NORTH
)

func (d *Direction) GetCoords(x1, y1 int) (x2, y2 int) {
	switch *d {
	case NORTH:
		x2, y2 = x1, y1-1
	case SOUTH:
		x2, y2 = x1, y1+1
	case WEST:
		x2, y2 = x1-1, y1
	case EAST:
		x2, y2 = x1+1, y1
	default:
		panic("Invalid Direction")
	}
	return
}

func (d Direction) Clockwise() (d2 Direction) {
	d2 = d + 1
	if d2 > 3 {
		d2 = 0
	}
	return
}

func (d Direction) CounterClockwise() (d2 Direction) {
	d2 = d - 1
	if d2 < 0 {
		d2 = 3
	}
	return
}
