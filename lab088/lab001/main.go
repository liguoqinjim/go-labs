package main

import (
	"log"
	"math"
)

type Point struct {
	X, Y float64
}

type Segment struct {
	P1, P2 Point
}

type Polygon struct {
	Sides []Segment
}

func Inside(p Point, polygon Polygon) (i bool) {
	for _, side := range polygon.Sides {
		if RayIntersectsSegment(p, side) {
			i = !i
		}
	}
	return
}

func RayIntersectsSegment(p Point, s Segment) bool {
	var a, b Point
	if s.P1.Y < s.P2.Y {
		a, b = s.P1, s.P2
	} else {
		a, b = s.P2, s.P1
	}

	for p.Y == a.Y || p.Y == b.Y {
		p.Y = math.Nextafter(p.Y, math.Inf(1))
	}
	for p.X == a.X || p.X == b.X {
		p.X = math.Nextafter(p.X, math.Inf(1))
	}

	if p.Y < a.Y || p.Y > b.Y {
		return false
	}

	if a.X > b.X {
		if p.X > a.X {
			return false
		}
		if p.X < b.X {
			return true
		}
	} else {
		if p.X > b.X {
			return false
		}
		if p.X < a.X {
			return true
		}
	}
	return (p.Y-a.Y)/(p.X-a.X) >= (b.Y-a.Y)/(b.X-a.X)
}

func main() {
	//测试RayIntersectsSegment
	pa := Point{X: 0, Y: 0}
	pb := Point{X: -2, Y: 4}
	px := Point{X: -1, Y: 2}
	s := Segment{P1: pa, P2: pb}
	if RayIntersectsSegment(px, s) {
		log.Println("相交")
	} else {
		log.Println("不相交")
	}
}
