package main

import (
	"github.com/kellydunn/golang-geo"
	"log"
)

func main() {
	p1 := geo.NewPoint(42.25, 120.2)
	p2 := geo.NewPoint(30.25, 112.2)

	//得到大圆距离
	dist := p1.GreatCircleDistance(p2)
	log.Printf("dist:%f", dist)

	//判断一个点是否在多边形中
	ps := make([]*geo.Point, 4)
	ps[0] = geo.NewPoint(-1, 1)
	ps[1] = geo.NewPoint(1, -1)
	ps[2] = geo.NewPoint(11, 9)
	ps[3] = geo.NewPoint(9, 11)
	p3 := geo.NewPoint(5, 5)
	p4 := geo.NewPoint(12, 12)
	polygon1 := geo.NewPolygon(ps)
	if polygon1.Contains(p3) {
		log.Println("多边形包含p3")
	} else {
		log.Println("多边形不包含p3")
	}
	if polygon1.Contains(p4) {
		log.Println("多边形包含p4")
	} else {
		log.Println("多边形不包含p4")
	}
}
