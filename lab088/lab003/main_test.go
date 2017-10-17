package main

import (
	"math"
	"testing"
)

var EPSILON float64 = 0.00000001

var cases = []struct {
	X     float64
	Y     float64
	Angle float64
}{
	{0, 0, 0},
	{0, 5, 90},
	{0, -5, 270},
	{5, 5, 45},
	{-5, 5, 135},
	{-5, -5, 225},
	{5, -5, 315},
	{math.Pow(3, 0.5), 1, 30},
	{-math.Pow(3, 0.5), 1, 150},
	{-math.Pow(3, 0.5), -1, 210},
	{math.Pow(3, 0.5), -1, 330},
}

func TestCalAngle(t *testing.T) {
	for _, c := range cases {
		angle := CalAngle(c.X, c.Y)
		if !floatEquals(angle, c.Angle) {
			t.Errorf("CalAngle(%f,%f) == %f, want %f", c.X, c.Y, angle, c.Angle)
		}
	}
}

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}
