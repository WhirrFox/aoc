package main

type Direction byte

const (
	UP    = Direction('^')
	DOWN  = Direction('v')
	LEFT  = Direction('<')
	RIGHT = Direction('>')
	NONE  = Direction('0')
)

func (d *Direction) GetCoords(x1, y1 int) (x2, y2 int) {
	switch *d {
	case UP:
		x2, y2 = x1, y1-1
	case DOWN:
		x2, y2 = x1, y1+1
	case LEFT:
		x2, y2 = x1-1, y1
	case RIGHT:
		x2, y2 = x1+1, y1
	default:
		panic("Invalid Direction")
	}
	return
}

type XY struct {
	X int
	Y int
}

type Robot struct {
	Pos         XY
	Moves       string
	currentMove int
}

func (r *Robot) Move(m *MapMap, d Direction) {
	x2, y2 := d.GetCoords(r.Pos.X, r.Pos.Y)
	switch m.Get(x2, y2) {
	case '#':
		return
	case '.':
		r.Pos.X, r.Pos.Y = x2, y2
	case 'O':
		if m.MoveBox(x2, y2, d) {
			r.Pos.X, r.Pos.Y = x2, y2
		}
	}
}

func (r *Robot) NextMove() Direction {
	r.currentMove++
	if r.currentMove > len(r.Moves) {
		return NONE
	}
	return Direction(r.Moves[r.currentMove-1])
}
