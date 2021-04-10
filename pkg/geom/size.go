package geom

import (
	"fmt"
	"github.com/maxfish/go-libs/pkg/imath"
)

type Size struct {
	W, H int
}

func SizeFromFloats(w, h float32) Size {
	return Size{int(w), int(h)}
}

func MaxSize(a,b Size) Size {
	return Size{
		W: imath.Max(a.W, b.W),
		H: imath.Max(a.H, b.H),
	}
}

func MinSize(a,b Size) Size {
	return Size{
		W: imath.Min(a.W, b.W),
		H: imath.Min(a.H, b.H),
	}
}

func (s Size) Sub(t Size) Size {
	return Size{
		W: s.W - t.W,
		H: s.H - t.H,
	}
}

func (s Size) Scale(factor float32) Size {
	return Size{
		W: int(float32(s.W) * factor),
		H: int(float32(s.H) * factor),
	}
}

func (s Size) ToString() string {
	return fmt.Sprintf("{w:%d,h:%d}", s.W, s.H)
}
