package main

type MapMap struct {
	M      []string
	M2     [][]int
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

func (m *MapMap) CreateMap2() {
	for _, line := range m.M {
		numbers := []int{}
		for _, r := range line {
			if r == '#' {
				numbers = append(numbers, -1)
			} else {
				numbers = append(numbers, 0)
			}
		}
		m.M2 = append(m.M2, numbers)
	}
}

func (m *MapMap) Spread(x1, y1 int, d Direction, score int) {
	m.Set2(x1, y1, score)
	if x1 == m.Width-2 && y1 == 1 {
		return
	}

	x2, y2 := d.GetCoords(x1, y1)
	if s := m.Get2(x2, y2); s == 0 || s > score+1 {
		m.Spread(x2, y2, d, score+1)
	}

	d2 := d.Clockwise()
	x2, y2 = d2.GetCoords(x1, y1)
	if s := m.Get2(x2, y2); s == 0 || s > score+1001 {
		m.Spread(x2, y2, d2, score+1001)
	}

	d3 := d.CounterClockwise()
	x2, y2 = d3.GetCoords(x1, y1)
	if s := m.Get2(x2, y2); s == 0 || s > score+1001 {
		m.Spread(x2, y2, d3, score+1001)
	}
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

func (m *MapMap) Set2(x, y, v int) {
	m.M2[y][x] = v
}

func (m *MapMap) Get2(x, y int) int {
	return m.M2[y][x]
}
