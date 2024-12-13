package main

type Area struct {
	Type      byte
	Area      int
	Perimeter int
}

func (a *Area) CalculateArea(m *Garden, x, y int) int {
	a.spread(m, x, y, '.')

	cat := Cat{
		X: x,
		Y: y - 1,
	}
	result := a.Area * cat.Meow(m)

	a.Type = '.'
	a.spread(m, x, y, 0)
	return result
}

func (a *Area) spread(m *Garden, x, y int, fillChar byte) {
	a.Area++
	Set(*m, x, y, fillChar)

	// Up
	if b := Get(*m, x, y-1); b == a.Type {
		a.spread(m, x, y-1, fillChar)
	} else if b != fillChar {
		a.Perimeter++
	}
	// Down
	if b := Get(*m, x, y+1); b == a.Type {
		a.spread(m, x, y+1, fillChar)
	} else if b != fillChar {
		a.Perimeter++
	}
	// Left
	if b := Get(*m, x-1, y); b == a.Type {
		a.spread(m, x-1, y, fillChar)
	} else if b != fillChar {
		a.Perimeter++
	}
	// Right
	if b := Get(*m, x+1, y); b == a.Type {
		a.spread(m, x+1, y, fillChar)
	} else if b != fillChar {
		a.Perimeter++
	}
}
