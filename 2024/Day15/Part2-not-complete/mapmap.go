package main

type MapMap struct {
	M      []string
	Width  int
	Height int
}

func NewMap(w, h int) MapMap {
	m := make([]string, h)
	var line string
	for range w {
		line += "."
	}
	for i := range h {
		m[i] = line
	}
	return MapMap{
		M:      m,
		Width:  w,
		Height: h,
	}
}

func (m *MapMap) MoveBox(x1, y1 int, d Direction) bool {
	x2, y2 := d.GetCoords(x1, y1)

	switch m.Get(x2, y2) {
	case '#':
		return false
	case '.':
		m.Set(x1, y1, '.')
		m.Set(x2, y2, 'O')
		return true
	case ']':
		return m.MoveBox(x2-1, y2, d)
	case '[':
		return m.moveBigBox(x2, y2, d)
	}
	return false
}

func (m *MapMap) Get(x, y int) byte {
	if x < 0 ||
		y < 0 ||
		x >= m.Width ||
		y >= m.Height {
		return 0
	}

	return m.M[y][x]
}

func (m *MapMap) Set(x, y int, b byte) {
	if x < 0 ||
		y < 0 ||
		x >= m.Width ||
		y >= m.Height {
		return
	}

	m.M[y] = replaceAtIndex(m.M[y], b, x)
}

func replaceAtIndex(in string, r byte, i int) string {
	out := []byte(in)
	out[i] = r
	return string(out)
}
