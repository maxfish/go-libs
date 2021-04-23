package geom

import (
	"github.com/maxfish/go-libs/pkg/imath"
)

// x,y are the current iteration coordinates.
// Should return true if the iteration needs to stop.
type BresenhamCallBack func(x int, y int) bool

// Iterate a line from x0,y0 to x1,y1.
// For each point on the line it calls the callBack function.
// Note: The order of the points is not preserved, and it's faster than IterateLineOrdered.
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

func IterateLineOrdered(x0, y0, x1, y1 int, callBack BresenhamCallBack) {
	deltaX := imath.Abs(x1 - x0)
	xStep := -1
	if x0 < x1 {
		xStep = 1
	}
	deltaY := imath.Abs(y1 - y0)
	yStep := -1
	if y0 < y1 {
		yStep = 1
	}

	if deltaX == 0 && deltaY == 0 {
		callBack(x0, y0)
		return
	}

	err := deltaX - deltaY
	err2 := 0

	for !(x0 == x1 && y0 == y1) {
		if callBack(x0, y0) {
			return
		}
		err2 = err * 2
		if err2 > -deltaY {
			err -= deltaY
			x0 += xStep
		}
		if err2 < deltaX {
			err += deltaX
			y0 += yStep
		}
	}
	// The last point on the line is not included by the loop
	callBack(x1, y1)
}
