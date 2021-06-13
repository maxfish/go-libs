package fgeom

import (
	"fmt"
	"github.com/maxfish/go-libs/pkg/fmath"
)

type Size struct {
	W, H float32
}

func SizeFromInt(w, h int) Size {
	return Size{float32(w), float32(h)}
}

func MaxSize(a,b Size) Size {
	return Size{
		W: fmath.Max(a.W, b.W),
		H: fmath.Max(a.H, b.H),
	}
}

func MinSize(a,b Size) Size {
	return Size{
		W: fmath.Min(a.W, b.W),
		H: fmath.Min(a.H, b.H),
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
		W: s.W * factor,
		H: s.H * factor,
	}
}

func (s Size) String() string {
	return fmt.Sprintf("{w:%.2f,h:%.2f}", s.W, s.H)
}
