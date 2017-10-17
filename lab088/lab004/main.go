package main

import "math"

func CalAngle(x1, y1, x2, y2 float64) float64 {
	var angle float64

	if x1 == x2 {
		if y1 == y2 {
			return 0
		} else if y2 > y1 {
			return 90
		} else {
			return 270
		}
	} else {
		k := (y2 - y1) / (x2 - x1)
		atan := math.Atan(k)
		angle = atan * 180 / math.Pi

		if x2 < x1 {
			angle += 180
		} else if x2 > x1 && y2 < y1 {
			angle += 360
		}
	}

	return angle
}
