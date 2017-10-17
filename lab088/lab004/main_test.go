package main

import (
	"math"
	"testing"
)

var EPSILON float64 = 0.00000001

var cases = []struct {
	X1, Y1 float64
	X2, Y2 float64
	Angle  float64
}{
	{2, 3, 2, 3, 0},
	{2, 3, 2, 5, 90},
	{2, 3, 2, -5, 270},
	{2, 3, 3, 4, 45},
	{2, 3, 1, 4, 135},
	{2, 3, 1, 2, 225},
	{2, 3, 3, 2, 315},
	{-3, -3, -3 + math.Pow(3, 0.5), -3 + 1, 30},
	{-3, -3, -3 - math.Pow(3, 0.5), -3 + 1, 150},
	{-3, -3, -3 - math.Pow(3, 0.5), -3 - 1, 210},
	{-3, -3, -3 + math.Pow(3, 0.5), -3 - 1, 330},
}

func TestCalAngle(t *testing.T) {
	for _, c := range cases {
		angle := CalAngle(c.X1, c.Y1, c.X2, c.Y2)
		if !floatEquals(angle, c.Angle) {
			t.Errorf("CalAngle(%f,%f,%f,%f) == %f, want %f", c.X1, c.Y1, c.X2, c.Y2, angle, c.Angle)
		}
	}
}

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}
