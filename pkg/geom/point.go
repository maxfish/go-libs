package geom

import "fmt"

type Point struct {
	X, Y int
}

func PointFromFloats(x,y float32) Point {
	return Point{X: int(x), Y: int(y)}
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func (p Point) EqualsTo(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) ToString() string {
	return fmt.Sprintf("{x:%d,y:%d}", p.X, p.Y)
}
