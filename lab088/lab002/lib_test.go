package lab002

import "testing"

var cases_kb = []struct {
	p1 *Point
	p2 *Point
	k  float64
	b  float64
}{
	{p1: &Point{0, 0}, p2: &Point{2, 2}, k: 1, b: 0},
	{p1: &Point{0, 0}, p2: &Point{-2, 2}, k: -1, b: 0},
	{p1: &Point{0, 1}, p2: &Point{1, 2}, k: 1, b: 1},
	{p1: &Point{0, 3}, p2: &Point{1, 3}, k: 0, b: 3},
	//{p1:&Point{5,0},&Point{5,5},}   //x=5这条线不调试
}

func TestCalKB(t *testing.T) {
	for _, c := range cases_kb {
		line := NewLine(c.p1, c.p2)
		line.CalKB()
		if line.K != c.k || line.B != c.b {
			t.Errorf("CalKB(%v,%v): k=%f,b=%f,want %f,%f", line.P1, line.P2, line.K, line.B, c.k, c.b)
		}
	}
}

var cases_point = []struct {
	p1       *Point
	p2       *Point
	p        *Point
	pointPos int
}{
	{p1: &Point{0, 0}, p2: &Point{2, 2}, p: &Point{1, 1}, pointPos: POINT_ON},
	{p1: &Point{0, 0}, p2: &Point{2, 2}, p: &Point{1, 0}, pointPos: POINT_UNDER},
	{p1: &Point{0, 0}, p2: &Point{2, 2}, p: &Point{1, 3}, pointPos: POINT_ABOVE},
	{p1: &Point{0, 0}, p2: &Point{-2, 2}, p: &Point{-1, 1}, pointPos: POINT_ON},
	{p1: &Point{0, 0}, p2: &Point{-2, 2}, p: &Point{-1, 0}, pointPos: POINT_UNDER},
	{p1: &Point{0, 0}, p2: &Point{-2, 2}, p: &Point{-1, 2}, pointPos: POINT_ABOVE},
}

func TestCalPointPos(t *testing.T) {
	for _, c := range cases_point {
		line := NewLine(c.p1, c.p2)
		line.CalKB()

		pointPos := line.CalPointPos(c.p)
		if pointPos != c.pointPos {
			t.Errorf("CalPointPos(%v,%v,%v): pointPos=%d,want %d", c.p1, c.p2, c.p, pointPos, c.pointPos)
		}
	}
}
