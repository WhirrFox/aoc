package main

type Direction int

const (
	NORTH = iota + 0
	EAST
	SOUTH
	WEST
)

type Guard struct {
	X         int
	Y         int
	Direction Direction
}

func NewGuard(x, y int) Guard {
	return Guard{
		X:         x,
		Y:         y,
		Direction: NORTH,
	}
}

func (g *Guard) Move() {
	switch g.Direction {
	case NORTH:
		g.Y--
	case EAST:
		g.X++
	case SOUTH:
		g.Y++
	case WEST:
		g.X--
	default:
		panic("Invalid Direction")
	}
}

func (g *Guard) Rotate() {
	g.Direction++
	if g.Direction > WEST {
		g.Direction = 0
	}
}

func (g *Guard) LookingAt() byte {
	switch g.Direction {
	case NORTH:
		return Get(g.X, g.Y-1)
	case EAST:
		return Get(g.X+1, g.Y)
	case SOUTH:
		return Get(g.X, g.Y+1)
	case WEST:
		return Get(g.X-1, g.Y)
	default:
		panic("Invalid direction")
	}
}
