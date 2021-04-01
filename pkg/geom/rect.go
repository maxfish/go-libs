package geom

import (
	"fmt"
	"github.com/maxfish/go-libs/pkg/math"
	"image"
)

type Rect struct {
	X, Y, W, H int
}

func RectFromArray(values [4]int) Rect {
	return Rect{values[0], values[1], values[2], values[3]}
}

func RectFromRectangle(r image.Rectangle) Rect {
	return Rect{r.Min.X, r.Min.Y, r.Dx(), r.Dy()}
}

func (r Rect) ToRectangle() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{r.X, r.Y},
		Max: image.Point{r.X + r.W, r.Y + r.H},
	}
}

func (r Rect) Left() int    { return r.X }
func (r Rect) Top() int     { return r.Y }
func (r Rect) Right() int   { return r.X + r.W }
func (r Rect) Bottom() int  { return r.Y + r.H }
func (r Rect) CenterX() int { return r.X + r.W/2 }
func (r Rect) CenterY() int { return r.Y + r.H/2 }

func (r Rect) Empty() bool {
	return r.X == 0 && r.Y == 0 && r.W == 0 && r.H == 0
}

func (r Rect) MoveTo(x, y int) Rect {
	r.X = x
	r.Y = y
	return r
}

func (r Rect) Translate(x, y int) Rect {
	r.X += x
	r.Y += y
	return r
}

func (r Rect) ResizeTo(w, h int) Rect {
	r.W = w
	r.H = h
	return r
}

func (r Rect) ShrinkByInsets(i Insets) Rect {
	return Rect{
		r.X + i.Left,
		r.Y + i.Top,
		r.W - i.Right - i.Left,
		r.H - i.Bottom - i.Top,
	}
}

func (r Rect) ShrinkByInt(i int) Rect {
	return Rect{
		r.X + i,
		r.Y + i,
		r.W - i*2,
		r.H - i*2,
	}
}

func (r Rect) CenterIn(o Rect) Rect {
	hW := (o.W - r.W) / 2
	hH := (o.H - r.H) / 2
	return Rect{
		r.X + hW,
		r.Y + hH,
		r.W + hW*2,
		r.H + hH*2,
	}
}

func (r Rect) AlignIn(b Rect, alignment Alignment) Rect {
	switch alignment.Horizontal {
	case AlignmentHLeft:
		r.X = b.X
	case AlignmentHCenter:
		r.X = b.X + (b.W-r.W)/2
	case AlignmentHRight:
		r.X = b.Right() - r.W
	}
	switch alignment.Vertical {
	case AlignmentVTop:
		r.Y = b.Y
	case AlignmentVCenter:
		r.Y = b.Y + (b.H-r.H)/2
	case AlignmentVBottom:
		r.Y = b.Bottom() - r.H
	}
	return r
}

func (r Rect) FitIn(b Rect, mode FitMode, alignment Alignment) Rect {
	switch mode {
	case FitModeFill:
		return b
	//case FitModeAlign:
	//	return r.AlignIn(b, alignment)
	case FitModeAspectFit:
		if r.W > r.H {
			f := float32(b.W) / float32(r.W)
			r.W = int(f * float32(r.W))
			r.H = int(f * float32(r.H))
		} else {
			f := float32(b.H) / float32(r.H)
			r.W = int(f * float32(r.W))
			r.H = int(f * float32(r.H))
		}
	case FitModeAspectFill:
		if r.W > r.H {
			f := float32(b.H) / float32(r.H)
			r.W = int(f * float32(r.W))
			r.H = int(f * float32(r.H))
		} else {
			f := float32(b.W) / float32(r.W)
			r.W = int(f * float32(r.W))
			r.H = int(f * float32(r.H))
		}
	}
	return r.AlignIn(b, alignment)
}

func (r Rect) UnionWith(other Rect) Rect {
	x1 := math.MinI(r.X, other.X)
	y1 := math.MinI(r.Y, other.Y)
	x2 := math.MaxI(r.Right(), other.Right())
	y2 := math.MaxI(r.Bottom(), other.Bottom())
	return Rect{X: x1, Y: y1, W: x2 - x1, H: y2 - y1}
}

func (r Rect) ContainsPoint(pointX, pointY int) bool {
	pointX -= r.X
	pointY -= r.Y
	return pointX >= 0 && pointY >= 0 && pointX < r.W && pointY < r.H
}

func (r Rect) Intersect(r2 Rect) bool {
	if r.X >= r2.X+r2.W || r2.X >= r.X+r.W {
		return false
	}
	if r.Y >= r2.Y+r2.H || r2.Y >= r.Y+r.H {
		return false
	}

	return true
}

// https://yal.cc/rectangle-circle-intersection-test/
func (r Rect) IntersectWithCircle(circleX, circleY, circleRadius int) bool {
	dX := circleX - math.MaxI(r.X, math.MinI(circleX, r.X+r.W))
	dY := circleY - math.MaxI(r.Y, math.MinI(circleY, r.Y+r.H))
	return (dX*dX + dY*dY) < (circleRadius * circleRadius)
}

func (r Rect) Intersection(s Rect) Rect {
	x2 := math.MinI(r.Right(), s.Right())
	y2 := math.MinI(r.Bottom(), s.Bottom())

	if r.X < s.X {
		r.X = s.X
	}
	if r.Y < s.Y {
		r.Y = s.Y
	}

	r.W = x2 - r.X
	r.H = y2 - r.Y

	if r.W > 0 && r.H > 0 {
		return r
	}
	return Rect{}
}

func (r Rect) EqualsTo(other Rect) bool {
	return r.X == other.X && r.Y == other.Y && r.W == other.W && r.H == other.H
}

func (r Rect) ToString() string {
	return fmt.Sprintf("{x:%d,y:%d,w:%d,h:%d}", r.X, r.Y, r.W, r.H)
}
