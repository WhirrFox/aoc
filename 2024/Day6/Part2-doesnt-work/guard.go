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

func (g *Guard) SetInFront(b byte) {
	switch g.Direction {
	case NORTH:
		Set(g.X, g.Y-1, b)
	case EAST:
		Set(g.X+1, g.Y, b)
	case SOUTH:
		Set(g.X, g.Y+1, b)
	case WEST:
		Set(g.X-1, g.Y, b)
	default:
		panic("Invalid direction")
	}
}

func (g *Guard) IsStuck() bool {
	stuckTimeout := 100000
	for stuckTimeout > 0 {
		stuckTimeout--
		if g.X < 0 || g.Y < 0 ||
			g.X >= width || g.Y >= height {
			return false
		}

		layoutMap[g.Y] = replaceAtIndex(layoutMap[g.Y], 'X', g.X)
		if g.LookingAt() == '#' {
			g.Rotate()
		} else {
			g.Move()
		}
	}
	return true
}
