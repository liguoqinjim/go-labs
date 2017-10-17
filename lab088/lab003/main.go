package main

import (
	"log"
	"math"
)

func main() {
	var x, y float64
	x, y = -5, 5

	angle := CalAngle(x, y)
	log.Println("angle:", angle)
}

//返回角度
func CalAngle(x, y float64) float64 {
	var angle float64

	if x == 0 {
		if y == 0 {
			angle = 0
		} else if y > 0 {
			angle = 90
		} else {
			angle = 270
		}
	} else {
		k := y / x
		atan := math.Atan(k)
		angle = atan * 180 / math.Pi

		if x < 0 {
			angle += 180
		} else if x > 0 && y < 0 {
			angle += 360
		}
	}

	return angle
}
