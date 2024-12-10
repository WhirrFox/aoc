package main

type Trailhead struct {
	X    int
	Y    int
	ends []Pos
}

func NewTrailhead(x, y int) Trailhead {
	return Trailhead{
		X: x,
		Y: y,
	}
}

func (t *Trailhead) CalculateScore(m *HikingMap) int {
	t.ends = []Pos{}
	t.MoveOneStep(m, t.X, t.Y, 0)
	return len(t.ends)
}

func (t *Trailhead) MoveOneStep(m *HikingMap, x, y, step int) {
	if step == 9 {
		for _, e := range t.ends {
			if e.X == x && e.Y == y {
				return
			}
		}
		t.ends = append(t.ends, Pos{x, y})
		return
	}

	nextStep := byte(step + 1)

	// Up
	if Get(*m, x, y-1)-'0' == nextStep {
		t.MoveOneStep(m, x, y-1, step+1)
	}
	// Down
	if Get(*m, x, y+1)-'0' == nextStep {
		t.MoveOneStep(m, x, y+1, step+1)
	}
	// Left
	if Get(*m, x-1, y)-'0' == nextStep {
		t.MoveOneStep(m, x-1, y, step+1)
	}
	// Right
	if Get(*m, x+1, y)-'0' == nextStep {
		t.MoveOneStep(m, x+1, y, step+1)
	}
}
