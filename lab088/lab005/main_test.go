package main

import "testing"

var handleAngleCases = []struct {
	Angle  float64
	Result float64
}{
	{0, 0},
	{45, 45},
	{405, 45},
	{-45, 315},
}

func TestHandleAngle(t *testing.T) {
	for _, c := range handleAngleCases {
		angle := HandleAngle(c.Angle)
		if angle != c.Result {
			t.Errorf("HandleAngle(%f) == %f, want %f", c.Angle, angle, c.Angle)
		}
	}
}

var sectorContainCases = []struct {
	Sector  *Sector
	X, Y    float64
	Contain bool
}{
	{&Sector{1, 1, 5, 5, 45, 5}, 4, 2, true},
	{&Sector{1, 1, 5, 5, 45, 5}, 2, 4, true},
	{&Sector{1, 1, 5, 5, 45, 5}, 0, 4, false},
	{&Sector{1, 1, 5, 5, 45, 5}, -2, -2, false},
	{&Sector{1, 1, 5, 5, 45, 5}, -2, -3, false},
	{&Sector{1, 1, 5, 5, 45, 5}, 6, 6, false},
	{&Sector{1, 0, 2, 1, 90, 5}, 4, -1, true},
	{&Sector{1, 0, 2, 1, 90, 5}, 4, 1, true},
	{&Sector{1, 0, 2, 1, 90, 5}, 0, 2, true},
	{&Sector{1, 0, 2, 1, 90, 5}, -1, 3, true},
	{&Sector{1, 0, 2, 1, 90, 5}, -1, 0, false},
	{&Sector{1, 0, 2, 1, 90, 5}, 0, -1, false},
	{&Sector{1, 0, 2, 1, 90, 5}, 1, -1, false},
	{&Sector{1, 0, 2, 1, 90, 5}, 3, -4, false},
}

func TestSector_ContainPoint(t *testing.T) {
	for _, c := range sectorContainCases {
		contain := c.Sector.ContainPoint(c.X, c.Y)
		if contain != c.Contain {
			t.Errorf("%v.ContainPoint(%f,%f) == %t, want %t", c.Sector, c.X, c.Y, contain, c.Contain)
		}
	}
}
