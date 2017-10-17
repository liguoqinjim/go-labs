package main

import "math"

func main() {

}

type Sector struct {
	X1, Y1     float64 //作为原点
	X2, Y2     float64 //和原点组成对称轴
	DeltaAngle float64 //左右偏移的角度
	R          float64 //扇形半径
}

func (sector *Sector) ContainPoint(x, y float64) bool {
	//判断是不是在园中
	d := math.Pow(math.Pow(x-sector.X1, 2)+math.Pow(y-sector.Y1, 2), 0.5)
	if d > sector.R {
		return false
	}

	//判断是否在角度内
	angleOrigin := CalAngle(sector.X1, sector.Y1, sector.X2, sector.Y2)
	angle1, angle2 := angleOrigin-sector.DeltaAngle, angleOrigin+sector.DeltaAngle
	angle1 = HandleAngle(angle1)
	angle2 = HandleAngle(angle2)

	angle := CalAngle(sector.X1, sector.Y1, x, y)

	if angle1 < angle2 {
		if angle < angle1 || angle > angle2 {
			return false
		}
	} else {
		if angle > angle2 && angle < angle1 {
			return false
		}
	}

	return true
}

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

//处理角度
func HandleAngle(angle float64) float64 {
	if angle >= 360 {
		angle -= 360
	}

	if angle < 0 {
		angle += 360
	}

	return angle
}
