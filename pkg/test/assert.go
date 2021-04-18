package test

import (
	"github.com/go-gl/mathgl/mgl32"
	"reflect"
	"testing"
)

const floatThreshold = 1e-6

func AssertEqual(t *testing.T, text string, expected interface{}, received interface{}) {
	if !reflect.DeepEqual(received, expected) {
		t.Errorf("%s failed\nexpected:\n%vreceived:\n%v", text, expected, received)
	}
}

func AssertMat4Equal(t *testing.T, text string, expected mgl32.Mat4, received mgl32.Mat4) {
	if !received.ApproxEqualThreshold(expected, floatThreshold) {
		t.Errorf("%s failed\nexpected:\n%vreceived:\n%v", text, expected, received)
	}
}

func AssertVec2Equal(t *testing.T, text string, expected mgl32.Vec2, received mgl32.Vec2) {
	if !received.ApproxEqualThreshold(expected, floatThreshold) {
		t.Errorf("%s failed\nexpected:\n%vreceived:\n%v", text, expected, received)
	}
}

func AssertVec3Equal(t *testing.T, text string, expected mgl32.Vec3, received mgl32.Vec3) {
	if !received.ApproxEqualThreshold(expected, floatThreshold) {
		t.Errorf("%s failed\nexpected:\n%vreceived:\n%v", text, expected, received)
	}
}
