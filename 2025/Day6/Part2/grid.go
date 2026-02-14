package main

type Grid [][]rune
type XY struct {
	x, y int
}

func (g *Grid) At(x, y int) rune {
	if x < 0 || x >= len((*g)[y]) || y < 0 || y >= len((*g)) {
		return ' '
	}
	return (*g)[y][x]
}

func (g *Grid) AtPos(pos XY) rune {
	return g.At(pos.x, pos.y)
}

func (g *Grid) Set(x, y int, s rune) {
	(*g)[y][x] = s
}

func (g *Grid) SetAtPos(pos XY, s rune) {
	g.Set(pos.x, pos.y, s)
}
