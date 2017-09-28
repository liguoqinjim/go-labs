package lab002

type Point struct {
	X, Y float64
}

type Line struct {
	P1, P2 *Point
	K, B   float64
}

func NewLine(p1, p2 *Point) *Line {
	return &Line{P1: p1, P2: p2}
}

//计算斜线的k和b
func (line *Line) CalKB() {
	line.CalK()
	line.CalB()
}

//计算k
func (line *Line) CalK() {
	if line.P1.X != line.P2.X {
		line.K = (line.P2.Y - line.P1.Y) / (line.P2.X - line.P1.X)
	}
}

//计算b
func (line *Line) CalB() {
	if line.P1.X != line.P2.X {
		line.B = line.P1.Y - line.P1.X*line.K
	}
}

//通过x计算y
func (line *Line) CalY(x float64) float64 {
	if line.P1.X != line.P2.X {
		return line.K*x + line.B
	} else {
		return 0
	}
}

const (
	POINT_ABOVE = 1  //在上方
	POINT_ON    = 0  //在线上
	POINT_UNDER = -1 //在下方
)

//判断点的位置，在斜线上方，下方和在线上
func (line *Line) CalPointPos(p *Point) int {
	if line.P1.X == line.P2.X {
		return POINT_ON
	}

	y := line.CalY(p.X)
	if y > p.Y {
		return POINT_UNDER
	} else if y < p.Y {
		return POINT_ABOVE
	} else {
		return POINT_ON
	}
}
