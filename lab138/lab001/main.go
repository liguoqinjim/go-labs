package main

import (
	"github.com/paulmach/go.geo"
	"log"
)

func main() {
	//库是链式调用的
	p := geo.NewPoint(0, 0)
	p.SetX(10).Add(geo.NewPoint(10, 10))
	log.Printf("p=%#v", p)

	//都是对指针的操作
	p1 := geo.NewPoint(5, 5)
	p2 := p1.SetX(20)
	if p1.Equals(p2) {
		log.Println("p1==p2")
	}

	//clone
	p2 = p1.Clone().SetX(30)
	if !p1.Equals(p2) {
		log.Println("p1!=p2")
	}

	//path和line的相交
	path := geo.NewPath()
	path.Push(geo.NewPoint(0, 0))
	path.Push(geo.NewPoint(1, 1))

	line := geo.NewLine(geo.NewPoint(0, 1), geo.NewPoint(1, 0))

	if path.Intersects(line) {
		points, segments := path.Intersection(line)

		for i := range points {
			log.Printf("Intersection %d at %v with path segment %d", i, points[i], segments[i][0])
		}
	}
}
