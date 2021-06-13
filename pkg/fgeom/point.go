package fgeom

import (
	"fmt"
	math "math"
)

var NullPoint = Point{X: -math.MaxFloat32, Y: -math.MaxFloat32}

type Point struct {
	X, Y float32
}

func PointFromInt(x, y int) Point {
	return Point{X: float32(x), Y: float32(y)}
}

func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p Point) Scale(scale float32) Point {
	return Point{X: p.X * scale, Y: p.Y * scale}
}

func (p Point) EqualsTo(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) String() string {
	return fmt.Sprintf("{x:%.2f,y:%.2f}", p.X, p.Y)
}
