package fgeom

import (
	"fmt"
	"github.com/maxfish/go-libs/pkg/fmath"
	"github.com/maxfish/go-libs/pkg/geom"
	"image"
)

type Rect struct {
	X, Y, W, H float32
}

func RectFromArray(values [4]float32) Rect {
	return Rect{X: values[0], Y: values[1], W: values[2], H: values[3]}
}

func RectFromInt(x, y, w, h int) Rect {
	return Rect{X: float32(x), Y: float32(y), W: float32(w), H: float32(h)}
}

func RectFromRectangle(r image.Rectangle) Rect {
	return Rect{X: float32(r.Min.X), Y: float32(r.Min.Y), W: float32(r.Dx()), H: float32(r.Dy())}
}

func (r Rect) ToRectangle() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: int(r.X), Y: int(r.Y)},
		Max: image.Point{X: int(r.X + r.W), Y: int(r.Y + r.H)},
	}
}

func (r Rect) Left() float32    { return r.X }
func (r Rect) Top() float32     { return r.Y }
func (r Rect) Right() float32   { return r.X + r.W }
func (r Rect) Bottom() float32  { return r.Y + r.H }
func (r Rect) CenterX() float32 { return r.X + r.W/2 }
func (r Rect) CenterY() float32 { return r.Y + r.H/2 }

func (r Rect) Center() Point {
	return Point{r.X + r.W/2, r.Y + r.H/2}
}

func (r Rect) MinPoint() Point {
	return Point{r.X, r.Y}
}

func (r Rect) MaxPoint() Point {
	return Point{r.Right(), r.Bottom()}
}

func (r Rect) Size() Size {
	return Size{r.W, r.H}
}

func (r Rect) Empty() bool {
	return r.X == 0 && r.Y == 0 && r.W == 0 && r.H == 0
}

func (r Rect) MoveTo(x, y float32) Rect {
	r.X = x
	r.Y = y
	return r
}

func (r Rect) Translate(x, y float32) Rect {
	r.X += x
	r.Y += y
	return r
}

func (r Rect) TranslatePoint(point Point) Rect {
	r.X += point.X
	r.Y += point.Y
	return r
}

func (r Rect) TranslateInt(x, y int) Rect {
	r.X += float32(x)
	r.Y += float32(y)
	return r
}

func (r Rect) ResizeTo(w, h float32) Rect {
	r.W = w
	r.H = h
	return r
}

func (r Rect) ShrinkByInsets(i Insets) Rect {
	return Rect{
		X: r.X + i.Left,
		Y: r.Y + i.Top,
		W: r.W - i.Right - i.Left,
		H: r.H - i.Bottom - i.Top,
	}
}

func (r Rect) ShrinkByInt(i float32) Rect {
	return Rect{
		X: r.X + i,
		Y: r.Y + i,
		W: r.W - i*2,
		H: r.H - i*2,
	}
}

func (r Rect) Scale(factor float32) Rect {
	return Rect{
		X: fmath.Round(r.X * factor),
		Y: fmath.Round(r.Y * factor),
		W: fmath.Round(r.W * factor),
		H: fmath.Round(r.H * factor),
	}
}

func (r Rect) CenterIn(o Rect) Rect {
	hW := (o.W - r.W) / 2
	hH := (o.H - r.H) / 2
	return Rect{
		X: r.X + hW,
		Y: r.Y + hH,
		W: r.W + hW*2,
		H: r.H + hH*2,
	}
}

func (r Rect) AlignIn(b Rect, alignment geom.Alignment) Rect {
	newRect := r
	if alignment&geom.AlignmentHLeft != 0 {
		newRect.X = b.X
	} else if alignment&geom.AlignmentHCenter != 0 {
		newRect.X = b.X + (b.W-r.W)/2
	} else if alignment&geom.AlignmentHRight != 0 {
		newRect.X = b.Right() - r.W
	}
	if alignment&geom.AlignmentVTop != 0 {
		newRect.Y = b.Y
	} else if alignment&geom.AlignmentVCenter != 0 {
		newRect.Y = b.Y + (b.H-r.H)/2
	} else if alignment&geom.AlignmentVBottom != 0 {
		newRect.Y = b.Bottom() - r.H
	}
	return newRect
}

func (r Rect) FitIn(b Rect, mode geom.FitMode, alignment geom.Alignment) Rect {
	switch mode {
	case geom.FitModeFill:
		return b
	//case FitModeAlign:
	//	return r.AlignIn(b, alignment)
	case geom.FitModeAspectFit:
		if r.W > r.H {
			f := b.W / r.W
			r.W = f * r.W
			r.H = f * r.H
		} else {
			f := b.H / r.H
			r.W = f * r.W
			r.H = f * r.H
		}
	case geom.FitModeAspectFill:
		if r.W > r.H {
			f := b.H / r.H
			r.W = f * r.W
			r.H = f * r.H
		} else {
			f := b.W / r.W
			r.W = f * r.W
			r.H = f * r.H
		}
	}
	return r.AlignIn(b, alignment)
}

func (r Rect) UnionWith(other Rect) Rect {
	x1 := fmath.Min(r.X, other.X)
	y1 := fmath.Min(r.Y, other.Y)
	x2 := fmath.Max(r.Right(), other.Right())
	y2 := fmath.Max(r.Bottom(), other.Bottom())
	return Rect{X: x1, Y: y1, W: x2 - x1, H: y2 - y1}
}

func (r Rect) ContainsPoint(pointX, pointY float32) bool {
	pointX -= r.X
	pointY -= r.Y
	return pointX >= 0 && pointY >= 0 && pointX < r.W && pointY < r.H
}

func (r Rect) IsContainedIn(b Rect) bool {
	return r.X >= b.X && r.Y >= b.Y && r.Right() <= b.Right() && r.Bottom() <= b.Bottom()
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
func (r Rect) IntersectWithCircle(circleX, circleY, circleRadius float32) bool {
	dX := circleX - fmath.Max(r.X, fmath.Min(circleX, r.X+r.W))
	dY := circleY - fmath.Max(r.Y, fmath.Min(circleY, r.Y+r.H))
	return (dX*dX + dY*dY) < (circleRadius * circleRadius)
}

func (r Rect) Intersection(s Rect) Rect {
	x2 := fmath.Min(r.Right(), s.Right())
	y2 := fmath.Min(r.Bottom(), s.Bottom())

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

func (r Rect) String() string {
	return fmt.Sprintf("{x:%.2f,y:%.2f,w:%.2f,h:%.2f}", r.X, r.Y, r.W, r.H)
}
