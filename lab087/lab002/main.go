package main

import (
	"github.com/kellydunn/golang-geo"
	"log"
)

func main() {
	//按照顺序的4个点
	var points1 []*geo.Point
	points1 = append(points1, geo.NewPoint(0, 0))
	points1 = append(points1, geo.NewPoint(5, 0))
	points1 = append(points1, geo.NewPoint(5, 5))
	points1 = append(points1, geo.NewPoint(0, 5))
	polygon1 := geo.NewPolygon(points1)
	point1 := geo.NewPoint(2, 2)
	point2 := geo.NewPoint(6, 6)
	if polygon1.Contains(point1) {
		log.Println("polygon1包含point1")
	} else {
		log.Println("polygon1不包含point1")
	}
	if polygon1.Contains(point2) {
		log.Println("polygon1包含point2")
	} else {
		log.Println("polygon1不包含point2")
	}

	//不按照顺序的4个点
	var points2 []*geo.Point
	points2 = append(points2, geo.NewPoint(0, 0))
	points2 = append(points2, geo.NewPoint(5, 5))
	points2 = append(points2, geo.NewPoint(5, 0))
	points2 = append(points2, geo.NewPoint(0, 5))
	polygon2 := geo.NewPolygon(points2)
	if polygon2.Contains(point1) {
		log.Println("polygon2包含point1")
	} else {
		log.Println("polygon2不包含point1")
	}
	if polygon2.Contains(point2) {
		log.Println("polygon2包含point2")
	} else {
		log.Println("polygon2不包含point2")
	}

	//用AddPoint方法来创建一个polygon
	polygon3 := geo.NewPolygon(nil)
	polygon3.Add(geo.NewPoint(0, 0))
	polygon3.Add(geo.NewPoint(5, 5))
	polygon3.Add(geo.NewPoint(5, 0))
	polygon3.Add(geo.NewPoint(0, 5))
	if polygon3.Contains(point1) {
		log.Println("polygon3包含point1")
	} else {
		log.Println("polygon3不包含point1")
	}
	if polygon3.Contains(point2) {
		log.Println("polygon3包含point2")
	} else {
		log.Println("polygon3不包含point2")
	}
}
