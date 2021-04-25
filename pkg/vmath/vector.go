package vmath

import (
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
