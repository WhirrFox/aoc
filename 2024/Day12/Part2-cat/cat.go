package main

type Direction int

const (
	RIGHT = iota + 0
	DOWN
	LEFT
	UP
)

// A cute lil cat who's walking on the outside of the areas to count the sides
type Cat struct {
	X         int
	Y         int
	Direction Direction
	sides     int
}

func (c *Cat) Meow(m *Garden) int {
	startX, startY := c.X, c.Y
	for {
		if c.LookingAt(m) == '.' {
			c.Rotate()
			c.Rotate()
			c.Rotate()
			c.sides++
		} else if c.LookingRight(m) != '.' {
			c.Rotate()
			c.sides++
			if c.LookingAt(m) != '.' {
				c.Move()
			}
		} else {
			c.Move()
		}

		if c.X == startX && c.Y == startY {
			return c.sides
		}
	}
}

func (c *Cat) Move() {
	switch c.Direction {
	case UP:
		c.Y--
	case RIGHT:
		c.X++
	case DOWN:
		c.Y++
	case LEFT:
		c.X--
	default:
		panic("Invalid Direction")
	}
}

func (c *Cat) Rotate() {
	c.Direction++
	if c.Direction > UP {
		c.Direction = 0
	}
}

func (c *Cat) LookingAt(m *Garden) byte {
	switch c.Direction {
	case UP:
		return Get(*m, c.X, c.Y-1)
	case RIGHT:
		return Get(*m, c.X+1, c.Y)
	case DOWN:
		return Get(*m, c.X, c.Y+1)
	case LEFT:
		return Get(*m, c.X-1, c.Y)
	default:
		panic("Invalid direction")
	}
}

func (c *Cat) LookingRight(m *Garden) byte {
	switch c.Direction {
	case UP:
		return Get(*m, c.X+1, c.Y)
	case RIGHT:
		return Get(*m, c.X, c.Y+1)
	case DOWN:
		return Get(*m, c.X-1, c.Y)
	case LEFT:
		return Get(*m, c.X, c.Y-1)
	default:
		panic("Invalid direction")
	}
}
