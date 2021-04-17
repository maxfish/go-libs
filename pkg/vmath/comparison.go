package vmath

import "github.com/maxfish/go-libs/pkg/fmath"
import "github.com/go-gl/mathgl/mgl32"

func ApproxEqualVec2(v1, v2 mgl32.Vec2, epsilon float32) bool {
	for i := range v1 {
		if !fmath.NearlyEqual(v1[i], v2[i], epsilon) {
			return false
		}
	}
	return true
}

func ApproxEqualVec3(v1, v2 mgl32.Vec3, epsilon float32) bool {
	for i := range v1 {
		if !fmath.NearlyEqual(v1[i], v2[i], epsilon) {
			return false
		}
	}
	return true
}
