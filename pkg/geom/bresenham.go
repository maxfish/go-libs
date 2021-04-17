package geom

import (
	"github.com/maxfish/go-libs/pkg/imath"
)

// x,y are the current iteration coordinates.
// Should return true if the iteration needs to stop.
type BresenhamCallBack func(x int, y int) bool

// Iterate a line from x0,y0 to x1,y1.
// For each point on the line it calls the callBack function.
func IterateLine(x0, y0, x1, y1 int, callBack BresenhamCallBack) {
	var swapCoords = imath.Abs(y1-y0) > imath.Abs(x1-x0)
	if swapCoords {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	deltaX := x1 - x0
	deltaY := imath.Abs(y1 - y0)
	err := deltaX >> 1
	y := y0

	yStep := -1
	if y0 < y1 {
		yStep = 1
	}

	if swapCoords {
		for x := x0; x <= x1; x++ {
			callBack(y, x)
			err -= deltaY
			if err < 0 {
				y += yStep
				err += deltaX
			}
		}
	} else {
		for x := x0; x <= x1; x++ {
			callBack(x, y)
			err -= deltaY
			if err < 0 {
				y += yStep
				err += deltaX
			}
		}
	}
}
