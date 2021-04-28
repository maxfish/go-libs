package vmath

import (
	"errors"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/maxfish/go-libs/pkg/fmath"
	"math"
)

func Vec2FromDegrees(angle float32) mgl32.Vec2 {
	angle *= fmath.DegreesToRadians
	cos := float32(math.Cos(float64(angle)))
	sin := float32(math.Sin(float64(angle)))
	return mgl32.Vec2{cos, sin}
}

func LerpVec2(a, b mgl32.Vec2, factor float32) mgl32.Vec2 {
	return mgl32.Vec2{
		(1-factor)*a[0] + factor*b[0],
		(1-factor)*a[1] + factor*b[1],
	}
}

func LerpVec4(a, b mgl32.Vec4, factor float32) mgl32.Vec4 {
	return mgl32.Vec4{
		(1-factor)*a[0] + factor*b[0],
		(1-factor)*a[1] + factor*b[1],
		(1-factor)*a[2] + factor*b[2],
		(1-factor)*a[3] + factor*b[3],
	}
}

// CircleToPolygon approximate a circle shape with a regular polygon
func CircleToPolygon(center mgl32.Vec2, radius float32, numSegments int, startAngle float32) ([]mgl32.Vec2, error) {
	if radius <= 0 {
		return nil, errors.New("radius cannot be <=0")
	}
	if numSegments < 3 {
		return nil, errors.New("numSegments must be >= 3")
	}
	point := mgl32.Rotate2D(startAngle).Mul2x1(mgl32.Vec2{radius, 0})
	vertices := make([]mgl32.Vec2, 0, numSegments*2)
	rotation := mgl32.Rotate2D((math.Pi * 2.0) / float32(numSegments))

	for index := 0; index < numSegments; index++ {
		p := point.Add(center)
		vertices = append(vertices, p)
		point = rotation.Mul2x1(point)
	}

	return vertices, nil
}

// BoundingBox returns the top left and the bottom right points of the 2D box bounding all the points passed.
func BoundingBox(points []mgl32.Vec2) (mgl32.Vec2, mgl32.Vec2) {
	var minX, minY, maxX, maxY float32
	minX = math.MaxFloat32
	minY = math.MaxFloat32
	maxX = -math.MaxFloat32
	maxY = -math.MaxFloat32
	for _, p := range points {
		if p.X() < minX {
			minX = p.X()
		}
		if p.X() > maxX {
			maxX = p.X()
		}
		if p.Y() < minY {
			minY = p.Y()
		}
		if p.Y() > maxY {
			maxY = p.Y()
		}
	}

	return mgl32.Vec2{minX, minY}, mgl32.Vec2{maxX, maxY}
}
